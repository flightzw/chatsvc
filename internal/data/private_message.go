package data

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/flightzw/chatsvc/internal/biz"
	"github.com/flightzw/chatsvc/internal/biz/query"
	"github.com/flightzw/chatsvc/internal/data/model"
	"github.com/flightzw/chatsvc/internal/enum"
	"github.com/flightzw/chatsvc/internal/utils/cache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gen"
)

type privateMessageRepo struct {
	*Data
	log *log.Helper
}

func NewPrivateMessageRepo(data *Data, logger log.Logger) biz.PrivateMessageRepo {
	return &privateMessageRepo{
		Data: data,
		log:  log.NewHelper(log.With(logger, "module", "chatsvc/data/PrivateMessageRepo")),
	}
}

func (repo *privateMessageRepo) CreatePrivateMessage(ctx context.Context, message *biz.PrivateMessage) (id int32, err error) {
	messageDo := repo.UseQuery(ctx).PrivateMessage.WithContext(ctx)
	data := &model.PrivateMessage{}

	if err = gconv.Struct(message, data); err != nil {
		return 0, errors.Wrap(err, "gconv.Struct")
	}
	if err = messageDo.Create(data); err != nil {
		return 0, errors.Wrap(err, "messageDo.Create")
	}
	// 今日新增消息数 + 1
	key := fmt.Sprintf(cache.RedisKeyTodayNewMessageCount, time.Now().Format("20060102"))
	if err = cache.IncrEX.Run(ctx, repo.redisClient, []string{key}, 84600).Err(); err != nil {
		return 0, errors.Wrap(err, "incrEX.Run")
	}
	return data.ID, nil
}

func (repo *privateMessageRepo) GetPrivateMessage(ctx context.Context, id int32) (data *biz.PrivateMessage, err error) {
	msg := repo.UseQuery(ctx).PrivateMessage
	messageDo := msg.WithContext(ctx)
	data = &biz.PrivateMessage{}

	message, err := messageDo.Where(msg.ID.Eq(id)).First()
	if err != nil {
		return nil, errors.Wrap(err, "messageDo.First")
	}
	if err = gconv.Struct(message, data); err != nil {
		return nil, errors.Wrap(err, "gconv.Struct")
	}
	return data, nil
}
func (repo *privateMessageRepo) ListPrivateMessage(ctx context.Context, queryFunc query.QueryFunc, page, pageSize int) (data []*biz.PrivateMessage, total int64, err error) {
	msg := repo.UseQuery(ctx).PrivateMessage
	messageDo := msg.WithContext(ctx)
	data = []*biz.PrivateMessage{}

	messageDo.DO = *queryFunc(&messageDo.DO).(*gen.DO)
	result, total, err := messageDo.FindByPage((page-1)*pageSize, pageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "messageDo.FindByPage")
	}
	if err = gconv.Structs(result, &data); err != nil {
		return nil, 0, errors.Wrap(err, "gconv.Structs")
	}
	return data, total, nil
}

func (repo *privateMessageRepo) UpdatePrivateMessageStatus(ctx context.Context, id int32, prev, curr enum.MessageStatus) (err error) {
	msg := repo.UseQuery(ctx).PrivateMessage
	messageDo := msg.WithContext(ctx).Omit(msg.ID, msg.CreatedAt)
	fmt.Println("UpdatePrivateMessageStatus:", prev, curr)
	_, err = messageDo.Where(msg.ID.Eq(id), msg.Status.Eq(int32(prev))).Update(msg.Status, curr)
	if err != nil {
		return errors.Wrap(err, "messageDo.Update")
	}
	return nil
}

func (repo *privateMessageRepo) ReadedPrivateMessage(ctx context.Context, userID, friendID int32) (err error) {
	msg := repo.UseQuery(ctx).PrivateMessage
	messageDo := msg.WithContext(ctx).Omit(msg.ID, msg.CreatedAt)

	_, err = messageDo.Where(msg.SendID.Eq(friendID), msg.RecvID.Eq(userID),
		msg.Status.Eq(enum.MessageStatusUnread)).Update(msg.Status, enum.MessageStatusReaded)
	if err != nil {
		return errors.Wrap(err, "messageDo.Update")
	}
	return nil
}

func (repo *privateMessageRepo) CountTodayNewMessageNum(ctx context.Context) (count int, err error) {
	result, err := repo.redisClient.Get(ctx, fmt.Sprintf(cache.RedisKeyTodayNewMessageCount, time.Now().Format("20060102"))).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, errors.Wrap(err, "redisClient.Get")
	}
	count, _ = strconv.Atoi(result)
	return count, nil
}

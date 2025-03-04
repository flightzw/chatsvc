package data

import (
	"context"

	"github.com/flightzw/chatsvc/internal/biz"
	"github.com/flightzw/chatsvc/internal/biz/query"
	"github.com/flightzw/chatsvc/internal/data/model"
	"github.com/flightzw/chatsvc/internal/entity"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"gorm.io/gen"
)

type friendRepo struct {
	*Data
	log *log.Helper
}

// NewFriendRepo .
func NewFriendRepo(data *Data, logger log.Logger) biz.FriendRepo {
	return &friendRepo{
		Data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *friendRepo) CreateFriend(ctx context.Context, friend *biz.Friend) (id int32, err error) {
	friendDo := repo.UseQuery(ctx).Friend.WithContext(ctx)
	data := &model.Friend{}

	if err = gconv.Struct(friend, data); err != nil {
		return 0, errors.Wrap(err, "gconv.Struct")
	}
	if err = friendDo.Create(data); err != nil {
		return 0, errors.Wrap(err, "friendDo.Create")
	}
	return data.ID, nil
}

func (repo *friendRepo) ListFriend(ctx context.Context, queryFunc query.QueryFunc, page, pageSize int) (data []*biz.Friend, total int64, err error) {
	friend := repo.UseQuery(ctx).Friend
	friendDo := friend.WithContext(ctx)
	data = []*biz.Friend{}

	friendDo.DO = *queryFunc(&friendDo.DO).(*gen.DO)
	result, total, err := friendDo.FindByPage((page-1)*pageSize, pageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "friendDo.FindByPage")
	}
	if err = gconv.Structs(result, &data); err != nil {
		return nil, 0, errors.Wrap(err, "gconv.Structs")
	}
	return data, total, nil
}

func (repo *friendRepo) UpdateFriend(ctx context.Context, id int32, data entity.AnyMap) (err error) {
	friend := repo.UseQuery(ctx).Friend
	friendDo := friend.WithContext(ctx).Omit(friend.ID, friend.CreatedAt)

	if _, err = friendDo.Where(friend.ID.Eq(id)).Updates(data.Assert()); err != nil {
		return errors.Wrap(err, "friendDo.Updates")
	}
	return nil
}

func (repo *friendRepo) DeleteFriend(ctx context.Context, id int32) (err error) {
	friend := repo.UseQuery(ctx).Friend
	friendDo := friend.WithContext(ctx)

	if _, err = friendDo.Where(friend.ID.Eq(id)).Delete(); err != nil {
		return errors.Wrap(err, "friendDo.Delete")
	}
	return nil
}

func (repo *friendRepo) GetFriendByFriendID(ctx context.Context, userID, friendID int32) (data *biz.Friend, err error) {
	friend := repo.UseQuery(ctx).Friend
	friendDo := friend.WithContext(ctx)
	data = &biz.Friend{}

	friendData, err := friendDo.Where(friend.UserID.Eq(userID), friend.FriendID.Eq(friendID)).First()
	if err != nil {
		return nil, errors.Wrap(err, "friendDo.First")
	}
	if err = gconv.Struct(friendData, data); err != nil {
		return nil, errors.Wrap(err, "gconv.Struct")
	}
	return data, nil
}

func (repo *friendRepo) DeleteFriendByFriendID(ctx context.Context, userID, friendID int32) (err error) {
	friend := repo.UseQuery(ctx).Friend
	friendDo := friend.WithContext(ctx)

	if _, err = friendDo.Where(friend.UserID.Eq(userID), friend.FriendID.Eq(friendID)).Delete(); err != nil {
		return errors.Wrap(err, "friendDo.Delete")
	}
	return nil
}

func (repo *friendRepo) UpdateFriendByFriendID(ctx context.Context, friendID int32, data entity.AnyMap) (err error) {
	friend := repo.UseQuery(ctx).Friend
	friendDo := friend.WithContext(ctx).Omit(friend.ID, friend.CreatedAt)

	if _, err = friendDo.Where(friend.FriendID.Eq(friendID)).Updates(data.Assert()); err != nil {
		return errors.Wrap(err, "friendDo.Updates")
	}
	return nil
}

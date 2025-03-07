package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/flightzw/chatsvc/api/chatsvc/errno"
	"github.com/flightzw/chatsvc/internal/biz/query"
	"github.com/flightzw/chatsvc/internal/conf"
	"github.com/flightzw/chatsvc/internal/entity"
	"github.com/flightzw/chatsvc/internal/enum"
	"github.com/flightzw/chatsvc/internal/utils/jwt"
	"github.com/flightzw/chatsvc/internal/vo"
	"github.com/flightzw/chatsvc/internal/ws"
	"github.com/flightzw/chatsvc/internal/ws/client"
)

type PrivateMessageQuery struct {
	Keyword     string `json:"keyword"`
	FriendID    int32  `json:"friend_id"`
	SendDateGte string `json:"send_date_gte"`
	SendDateLte string `json:"send_date_lte"`
}

type PrivateMessage struct {
	ID        int32              `json:"id"`
	SendID    int32              `json:"send_id"`    // 用户 uid
	RecvID    int32              `json:"recv_id"`    // 好友 uid
	Content   string             `json:"content"`    // 发送内容
	Type      enum.MessageType   `json:"type"`       // 消息类型
	Status    enum.MessageStatus `json:"status"`     // 状态 0:未送达 1:已送达 2:撤回 3:已读
	CreatedAt *gtime.Time        `json:"created_at"` // 创建时间
}

type PrivateMessageRepo interface {
	CreatePrivateMessage(context.Context, *PrivateMessage) (int32, error)
	GetPrivateMessage(context.Context, int32) (*PrivateMessage, error)
	ListPrivateMessage(ctx context.Context, queryFunc query.QueryFunc, page, pageSize int) ([]*PrivateMessage, int64, error)
	UpdatePrivateMessageStatus(ctx context.Context, id int32, prev, curr enum.MessageStatus) error
	ReadedPrivateMessage(ctx context.Context, userID, friendID int32) error

	CountTodayNewMessageNum(ctx context.Context) (count int, err error)
}

type PrivateMessageUsecase struct {
	repo       PrivateMessageRepo
	friendRepo FriendRepo
	userRepo   UserRepo

	log        *log.Helper
	conf       *conf.Server
	chatClient *client.ChatClient
}

func NewPrivateMessageUsecase(repo PrivateMessageRepo, friendRepo FriendRepo, userRepo UserRepo, conf *conf.Server,
	logger log.Logger, chatClient *client.ChatClient) *PrivateMessageUsecase {
	return &PrivateMessageUsecase{
		repo:       repo,
		friendRepo: friendRepo,
		userRepo:   userRepo,
		log:        log.NewHelper(log.With(logger, "module", "chatsvc/biz/PrivateMessageUsecase")),
		chatClient: chatClient,
	}
}

func (uc *PrivateMessageUsecase) SendPrivateMessage(ctx context.Context, message *PrivateMessage) (data *vo.PrivateMessageVO, err error) {
	count, err := uc.repo.CountTodayNewMessageNum(ctx)
	if err != nil {
		return nil, errno.ErrorSystemInternalFailure("此功能暂不可用").WithCause(err)
	}
	if count >= int(uc.conf.Limit.DailyMaxNewMsgNum) {
		return nil, errno.ErrorDeniedAccess("今日消息发送数已达上限")
	}

	message.SendID, _ = jwt.GetUserInfo(ctx)
	if _, err := uc.friendRepo.GetFriendByFriendID(ctx, message.SendID, message.RecvID); err != nil {
		return nil, errno.ErrorParamInvalid("您不是对方好友，无法发送消息").WithCause(err)
	}

	message.CreatedAt = gtime.Now()
	id, err := uc.repo.CreatePrivateMessage(ctx, message)
	if err != nil {
		return nil, errno.ErrorDataSaveFailed("暂存消息时出错").WithCause(err)
	}

	result := newPrivateMessageVO(message)
	result.ID = id

	// 发送给好友
	if err = uc.sendPrivateMessage(ctx, message.RecvID, result); err != nil {
		return nil, errno.ErrorMessageSendFailed("发送消息时出错").WithCause(err)
	}
	// 响应结果
	return result, nil
}

// 处理消息结果通知
func (uc *PrivateMessageUsecase) ProcessMessageSendResult(ctx context.Context) {
	var (
		err        error
		notifyChan = uc.chatClient.GetResultNotifyChan(ctx)
	)
	for data := range notifyChan {
		if !data.Success {
			continue
		}
		msg := &vo.PrivateMessageVO{}
		if err = gconv.Struct(data.Data.Data, msg); err != nil {
			uc.log.Error("gconv.Struct:", err)
		}
		if msg.Status != enum.MessageStatusUnsend {
			fmt.Println("[svc] message sended, no need update.")
			continue
		}
		err = uc.repo.UpdatePrivateMessageStatus(ctx, msg.ID, enum.MessageStatusUnsend, enum.MessageStatusUnread)
		if err != nil {
			uc.log.Error("repo.UpdatePrivateMessageStatus:", err)
		}
	}
}

func (uc *PrivateMessageUsecase) RecallPrivateMessage(ctx context.Context, id int32) (err error) {
	message, err := uc.repo.GetPrivateMessage(ctx, id)
	if err != nil {
		return errno.ErrorDataNotFound("消息不存在").WithCause(err)
	}
	userID, _ := jwt.GetUserInfo(ctx)
	if message.SendID != userID {
		return errno.ErrorParamInvalid("你不能撤回这条消息")
	}
	if message.Status == enum.MessageStatusRecall {
		return nil
	}
	if message.CreatedAt.Before(gtime.Now().Add(-15 * time.Minute)) {
		return errno.ErrorParamInvalid("消息发送已超过%d分钟，无法撤回", 15)
	}

	err = uc.repo.UpdatePrivateMessageStatus(ctx, id, message.Status, enum.MessageStatusRecall)
	if err != nil {
		return errno.ErrorDataUpdateFailed("撤回消息时出错").WithCause(err)
	}

	result := newPrivateMessageVO(message)
	result.Status = enum.MessageStatusRecall
	result.Content = "你撤回了一条消息"
	uc.sendPrivateMessage(ctx, message.SendID, result)
	result.Content = "对方撤回了一条消息"
	uc.sendPrivateMessage(ctx, message.RecvID, result)
	return
}

// 拉取离线消息
func (uc *PrivateMessageUsecase) PullOfflinePrivateMessage(ctx context.Context, id int32) (err error) {
	userID, _ := jwt.GetUserInfo(ctx)

	friends, _, err := uc.friendRepo.ListFriend(ctx, func(do query.QueryChain) query.QueryChain {
		return do.Where(query.NewFriendQuery().UserID.Eq(userID))
	}, 1, 100)
	if err != nil {
		return errno.ErrorDataQueryFailed("获取好友列表数据时出错").WithCause(err)
	}
	if len(friends) == 0 {
		return nil
	}

	friendIds := []int32{}
	for _, friend := range friends {
		friendIds = append(friendIds, friend.FriendID)
	}
	qf := func(do query.QueryChain) query.QueryChain {
		msg := query.NewPrivateMessageQuery()
		return do.Where(msg.ID.Gt(id)).
			Where(msg.Status.Neq(int32(enum.MessageStatusRecall))).
			Where(do.Where(do.Where(msg.SendID.Eq(userID), msg.RecvID.In(friendIds...))).
				Or(msg.RecvID.Eq(userID), msg.SendID.In(friendIds...)),
			).
			Where(msg.CreatedAt.Gt(time.Now().AddDate(0, 0, -30)))
	}
	data, _, err := uc.repo.ListPrivateMessage(ctx, qf, 1, 1000)
	if err != nil {
		return errno.ErrorDataQueryFailed("拉取离线消息时出错").WithCause(err)
	}
	for _, msg := range data {
		uc.sendPrivateMessage(ctx, userID, newPrivateMessageVO(msg))
	}
	uc.chatClient.SendMessage(ctx, &ws.MessageWrapper{
		RecvIds: []int32{userID},
		Data:    &ws.SendMessage{Action: enum.ActionTypeOfflinePush, Data: false},
	})
	return nil
}

func (uc *PrivateMessageUsecase) ReadedPrivateMessage(ctx context.Context, friendID int32) (err error) {
	userID, _ := jwt.GetUserInfo(ctx)

	// 通知用户自己消息已读
	err = uc.chatClient.SendMessage(ctx, &ws.MessageWrapper{
		RecvIds: []int32{userID},
		Data: &ws.SendMessage{Action: enum.ActionTypeMessageReaded, Data: entity.AnyMap{
			"chat_id": friendID, "self_readed": true,
		}},
	})
	if err != nil {
		return errno.ErrorMessageSendFailed("更新消息状态时出错").WithCause(err)
	}
	// 通知好友消息已被查看
	err = uc.chatClient.SendMessage(ctx, &ws.MessageWrapper{
		RecvIds: []int32{friendID},
		Data: &ws.SendMessage{Action: enum.ActionTypeMessageReaded, Data: entity.AnyMap{
			"chat_id": userID, "self_readed": false,
		}},
	})
	if err != nil {
		return errno.ErrorMessageSendFailed("更新消息状态时出错").WithCause(err)
	}
	if err = uc.repo.ReadedPrivateMessage(ctx, userID, friendID); err != nil {
		return errno.ErrorDataUpdateFailed("更新消息状态时出错").WithCause(err)
	}
	return nil
}

func (uc *PrivateMessageUsecase) ListPrivateMessage(ctx context.Context, params *PrivateMessageQuery, page, pageSize int32) (data []*vo.PrivateMessageVO, total int64, err error) {
	userID, _ := jwt.GetUserInfo(ctx)
	qf := func(do query.QueryChain) query.QueryChain {
		msg := query.NewPrivateMessageQuery()
		cpDo := do
		if params.Keyword != "" {
			cpDo = cpDo.Where(msg.Content.Like("%" + params.Keyword + "%"))
		}
		if params.SendDateGte != "" {
			cpDo = cpDo.Where(msg.CreatedAt.Gte(gtime.New(params.SendDateGte).Time))
		}
		if params.SendDateLte != "" {
			cpDo = cpDo.Where(msg.CreatedAt.Lte(gtime.New(params.SendDateLte).Time))
		}
		return cpDo.Where(msg.Status.Neq(int32(enum.MessageStatusRecall))).
			Where(do.Where(do.Where(msg.SendID.Eq(userID), msg.RecvID.Eq(params.FriendID))).
				Or(msg.RecvID.Eq(userID), msg.SendID.Eq(params.FriendID)),
			).
			Where(msg.CreatedAt.Gt(time.Now().AddDate(0, 0, -180))).Order(msg.ID.Desc())
	}
	msgs, total, err := uc.repo.ListPrivateMessage(ctx, qf, int(page), int(pageSize))
	if err != nil {
		return nil, 0, errno.ErrorDataQueryFailed("拉取离线消息时出错").WithCause(err)
	}
	data = make([]*vo.PrivateMessageVO, 0, len(msgs))
	for _, msg := range msgs {
		data = append(data, newPrivateMessageVO(msg))
	}
	return data, total, nil
}

func (uc *PrivateMessageUsecase) sendPrivateMessage(ctx context.Context, recvId int32, data *vo.PrivateMessageVO) error {
	return uc.chatClient.SendMessage(ctx, &ws.MessageWrapper{
		RecvIds: []int32{recvId},
		Data: &ws.SendMessage{
			Action: enum.ActionTypePrivateMessage,
			Data:   data,
		},
		NotifyResult: true,
	})
}

func newPrivateMessageVO(data *PrivateMessage) *vo.PrivateMessageVO {
	return &vo.PrivateMessageVO{
		ID:        data.ID,
		SendID:    data.SendID,
		RecvID:    data.RecvID,
		Content:   data.Content,
		Type:      data.Type,
		Status:    data.Status,
		CreatedAt: data.CreatedAt,
	}
}

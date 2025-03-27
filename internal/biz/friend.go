package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/flightzw/chatsvc/api/chatsvc/errno"
	"github.com/flightzw/chatsvc/internal/biz/query"
	"github.com/flightzw/chatsvc/internal/entity"
	"github.com/flightzw/chatsvc/internal/enum"
	"github.com/flightzw/chatsvc/internal/utils/jwt"
	"github.com/flightzw/chatsvc/internal/vo"
	"github.com/flightzw/chatsvc/internal/ws/client"
)

type Friend struct {
	ID              int32       `json:"id"`
	UserID          int32       `json:"user_id"`           // 用户id
	FriendID        int32       `json:"friend_id"`         // 好友id
	FriendNickname  string      `json:"friend_nickname"`   // 昵称
	FriendAvatarURL string      `json:"friend_avatar_url"` // 头像url
	Remark          string      `json:"remark"`            // 备注
	CreatedAt       *gtime.Time `json:"created_at"`        // 创建时间
}

type FriendRepo interface {
	TransactionInterface

	CreateFriend(ctx context.Context, friend *Friend) (id int32, err error)
	ListFriend(ctx context.Context, qf query.QueryFunc, page, pageSize int) (data []*Friend, total int64, err error)
	UpdateFriend(ctx context.Context, id int32, data entity.AnyMap) (err error)
	DeleteFriend(ctx context.Context, id int32) (err error)

	GetFriendByFriendID(ctx context.Context, userID, friendID int32) (data *Friend, err error)
	DeleteFriendByFriendID(ctx context.Context, userID, friendID int32) (err error)
	UpdateFriendByFriendID(ctx context.Context, friendID int32, data entity.AnyMap) (err error)
}

type FriendUsecase struct {
	repo     FriendRepo
	userRepo UserRepo

	log        *log.Helper
	chatClient *client.ChatClient
}

func NewFriendUsecase(repo FriendRepo, userRepo UserRepo, logger log.Logger, chatClient *client.ChatClient) *FriendUsecase {
	return &FriendUsecase{
		repo:       repo,
		userRepo:   userRepo,
		log:        log.NewHelper(log.With(logger, "module", "chatsvc/biz/FriendUsecase")),
		chatClient: chatClient,
	}
}

func (uc *FriendUsecase) AddFriend(ctx context.Context, userID int32) (err error) {
	id, _ := jwt.GetUserInfo(ctx)
	if id == userID {
		return errno.ErrorParamInvalid("不能将自己添加为好友")
	}
	if _, err = uc.repo.GetFriendByFriendID(ctx, id, userID); err == nil {
		return nil
	}

	if err = uc.bindFriend(ctx, int32(id), userID); err != nil {
		return err
	}
	if err = uc.bindFriend(ctx, userID, int32(id)); err != nil {
		return err
	}
	return nil
}

func (uc *FriendUsecase) bindFriend(ctx context.Context, userID, friendID int32) (err error) {
	friend, err := uc.userRepo.GetUser(ctx, friendID)
	if err != nil {
		return errno.ErrorUserNotFound("未找到指定的用户信息").WithCause(err)
	}
	_, err = uc.repo.CreateFriend(ctx, &Friend{
		UserID:          userID,
		FriendID:        friend.ID,
		FriendNickname:  friend.Nickname,
		FriendAvatarURL: friend.AvatarURL,
	})
	if err != nil {
		return errno.ErrorDataSaveFailed("添加好友时出错").WithCause(err)
	}
	return nil
}

func (uc *FriendUsecase) GetFriend(ctx context.Context, friendID int32) (data *vo.FriendVO, err error) {
	userID, _ := jwt.GetUserInfo(ctx)
	friend, err := uc.repo.GetFriendByFriendID(ctx, userID, friendID)
	if err != nil {
		return nil, errno.ErrorDataQueryFailed("未找到好友信息").WithCause(err)
	}
	user, err := uc.userRepo.GetUser(ctx, friend.FriendID)
	if err != nil {
		return nil, errno.ErrorDataNotFound("未找到好友信息").WithCause(err)
	}
	onlineMap := uc.chatClient.IsOnline(ctx, friendID)
	return newFriendVO(user, friend.Remark, onlineMap[friendID]), nil
}

func (uc *FriendUsecase) ListFriend(ctx context.Context) (data []*vo.FriendVO, err error) {
	id, _ := jwt.GetUserInfo(ctx)
	querys := func(do query.QueryChain) query.QueryChain {
		return do.Where(query.NewFriendQuery().UserID.Eq(id))
	}
	result, _, err := uc.repo.ListFriend(ctx, querys, 1, 100)
	if err != nil {
		return nil, errno.ErrorDataQueryFailed("获取好友列表数据时出错").WithCause(err)
	}

	friendIds := []int32{}
	for _, friend := range result {
		friendIds = append(friendIds, friend.FriendID)
	}
	userMap, err := uc.getUserMap(ctx, friendIds)
	if err != nil {
		return nil, err
	}
	onlineMap := uc.chatClient.IsOnline(ctx, friendIds...)
	for _, friend := range result {
		data = append(data, newFriendVO(userMap[friend.FriendID], friend.Remark, onlineMap[friend.FriendID]))
	}
	return data, nil
}

// 更新好友备注
func (uc *FriendUsecase) UpdateFriend(ctx context.Context, friend *Friend) (err error) {
	id, _ := jwt.GetUserInfo(ctx)
	data, err := uc.repo.GetFriendByFriendID(ctx, id, friend.FriendID)
	if err != nil {
		return errno.ErrorDataQueryFailed("未找到好友信息").WithCause(err)
	}
	if err = uc.repo.UpdateFriend(ctx, data.ID, entity.AnyMap{"remark": friend.Remark}); err != nil {
		return errno.ErrorDataUpdateFailed("更新好友信息时出错").WithCause(err)
	}
	return nil
}

func (uc *FriendUsecase) RemoveFriend(ctx context.Context, friendID int32) (err error) {
	id, _ := jwt.GetUserInfo(ctx)
	if err = uc.repo.DeleteFriendByFriendID(ctx, id, friendID); err != nil {
		return errno.ErrorDataRemoveFailed("删除好友信息时出错").WithCause(err)
	}
	if err = uc.repo.DeleteFriendByFriendID(ctx, friendID, id); err != nil {
		return errno.ErrorDataRemoveFailed("删除好友信息时出错").WithCause(err)
	}
	return nil
}

func (uc *FriendUsecase) getUserMap(ctx context.Context, ids []int32) (map[int32]*User, error) {
	qf := func(do query.QueryChain) query.QueryChain {
		return do.Where(query.NewUserQuery().ID.In(ids...))
	}
	data, _, err := uc.userRepo.ListUser(ctx, qf, 1, len(ids))
	if err != nil {
		return nil, errno.ErrorDataQueryFailed("获取好友信息时出错").WithCause(err)
	}
	userMap := map[int32]*User{}
	for _, user := range data {
		userMap[user.ID] = user
	}
	if len(userMap) == len(ids) {
		return userMap, nil
	}
	for _, userID := range ids {
		if _, ok := userMap[userID]; !ok {
			userMap[userID] = &User{ID: userID, Username: "用户已注销", Nickname: "用户已注销"}
		}
	}
	return userMap, nil
}

func newFriendVO(user *User, remark string, isOnline bool) *vo.FriendVO {
	if user.Type == enum.UserTypeAI {
		isOnline = true
	}
	return &vo.FriendVO{
		ID:        user.ID,
		Type:      user.Type,
		Username:  user.Username,
		Nickname:  user.Nickname,
		AvatarUrl: user.AvatarURL,
		Gender:    user.Gender,
		Signature: user.Signature,
		Remark:    remark,
		IsOnline:  isOnline,
	}
}

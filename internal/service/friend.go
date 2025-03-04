package service

import (
	"context"

	"github.com/flightzw/chatsvc/api/chatsvc/errno"
	v1 "github.com/flightzw/chatsvc/api/chatsvc/v1"
	"github.com/flightzw/chatsvc/internal/biz"
	"github.com/gogf/gf/v2/util/gconv"
)

type FriendService struct {
	v1.UnimplementedFriendServiceServer

	uc *biz.FriendUsecase
}

func NewFriendService(uc *biz.FriendUsecase) *FriendService {
	return &FriendService{uc: uc}
}

// 添加好友
func (svc *FriendService) AddFriend(ctx context.Context, req *v1.AddFriendRequest) (*v1.AddFriendReply, error) {
	if err := svc.uc.AddFriend(ctx, req.UserId); err != nil {
		return nil, err
	}
	return &v1.AddFriendReply{}, nil
}
func (svc *FriendService) GetFriend(ctx context.Context, req *v1.GetFriendRequest) (*v1.GetFriendReply, error) {
	data, err := svc.uc.GetFriend(ctx, req.FriendId)
	if err != nil {
		return nil, err
	}
	reply := &v1.GetFriendReply{}
	if err := gconv.Struct(data, &reply.Data); err != nil {
		return nil, errno.ErrorVoConvertFailed("数据转换时出错: %v", err)
	}
	return reply, nil
}

// 好友列表
func (svc *FriendService) ListFriend(ctx context.Context, req *v1.ListFriendRequest) (*v1.ListFriendReply, error) {
	data, err := svc.uc.ListFriend(ctx)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListFriendReply{}
	if err := gconv.Struct(data, &reply.Data); err != nil {
		return nil, errno.ErrorVoConvertFailed("数据转换时出错: %v", err)
	}
	return reply, nil
}

// 更新好友备注
func (svc *FriendService) UpdateFriend(ctx context.Context, req *v1.UpdateFriendRequest) (*v1.UpdateFriendReply, error) {
	data := &biz.Friend{}
	if err := gconv.Struct(req, data); err != nil {
		return nil, errno.ErrorDoConvertFailed("数据转换时出错: %v", err)
	}
	if err := svc.uc.UpdateFriend(ctx, data); err != nil {
		return nil, err
	}
	return &v1.UpdateFriendReply{}, nil
}

// 移除好友
func (svc *FriendService) RemoveFriend(ctx context.Context, req *v1.RemoveFriendRequest) (*v1.RemoveFriendReply, error) {
	if err := svc.uc.RemoveFriend(ctx, req.FriendId); err != nil {
		return nil, err
	}
	return &v1.RemoveFriendReply{}, nil
}

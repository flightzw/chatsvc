package service

import (
	"context"

	"github.com/flightzw/chatsvc/api/chatsvc/errno"
	v1 "github.com/flightzw/chatsvc/api/chatsvc/v1"
	"github.com/flightzw/chatsvc/internal/biz"
	"github.com/gogf/gf/v2/util/gconv"
)

// svc *UserService v1.UserServiceServer
type UserService struct {
	v1.UnimplementedUserServiceServer

	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// 获取当前用户信息
func (svc *UserService) GetUserSelf(ctx context.Context, _ *v1.GetUserSelfRequest) (*v1.GetUserSelfReply, error) {
	data, err := svc.uc.GetUserSelf(ctx)
	if err != nil {
		return nil, err
	}
	reply := &v1.GetUserSelfReply{}
	if err := gconv.Struct(data, &reply.Data); err != nil {
		return nil, errno.ErrorVoConvertFailed("数据转换时出错").WithCause(err)
	}
	return reply, nil
}

// 变更用户信息
func (svc *UserService) UpdateUserInfo(ctx context.Context, req *v1.UpdateUserInfoRequest) (*v1.UpdateUserInfoReply, error) {
	data := &biz.User{}
	if err := gconv.Struct(req, data); err != nil {
		return nil, errno.ErrorDoConvertFailed("数据转换时出错").WithCause(err)
	}
	if err := svc.uc.UpdateUserInfo(ctx, data); err != nil {
		return nil, err
	}
	return &v1.UpdateUserInfoReply{}, nil
}

// 获取用户信息
func (svc *UserService) GetUserByID(ctx context.Context, req *v1.GetUserByIDRequest) (*v1.GetUserByIDReply, error) {
	data, err := svc.uc.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	reply := &v1.GetUserByIDReply{}
	if err := gconv.Struct(data, &reply.Data); err != nil {
		return nil, errno.ErrorVoConvertFailed("数据转换时出错").WithCause(err)
	}
	return reply, nil
}

// 用户列表
func (svc *UserService) ListUserInfo(ctx context.Context, req *v1.ListUserInfoRequest) (*v1.ListUserInfoReply, error) {
	data, err := svc.uc.ListUserInfo(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListUserInfoReply{}
	if err = gconv.Struct(data, &reply.Data); err != nil {
		return nil, errno.ErrorVoConvertFailed("数据转换时出错").WithCause(err)
	}
	return reply, nil
}

// 更改密码
func (svc *UserService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*v1.UpdatePasswordReply, error) {
	err := svc.uc.UpdatePassword(ctx, req.OldPassword, req.NewPassword)
	if err != nil {
		return nil, err
	}
	return &v1.UpdatePasswordReply{}, nil
}

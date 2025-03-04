package service

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/flightzw/chatsvc/api/chatsvc/errno"
	v1 "github.com/flightzw/chatsvc/api/chatsvc/v1"
	"github.com/flightzw/chatsvc/internal/biz"
)

type LoginService struct {
	v1.UnimplementedLoginServiceServer

	uc *biz.UserUsecase
}

func NewLoginService(uc *biz.UserUsecase) *LoginService {
	return &LoginService{uc: uc}
}

func (svc *LoginService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	user := &biz.User{}
	if err := gconv.Struct(req, user); err != nil {
		return nil, errno.ErrorDoConvertFailed("数据转换时出错").WithCause(err)
	}
	err := svc.uc.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{}, nil
}

func (svc *LoginService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	user := &biz.User{}
	if err := gconv.Struct(req, user); err != nil {
		return nil, errno.ErrorDoConvertFailed("数据转换时出错").WithCause(err)
	}
	data, err := svc.uc.Login(ctx, user, req.RememberMe)
	if err != nil {
		return nil, err
	}
	reply := &v1.LoginReply{}
	if err := gconv.Struct(data, &reply.Data); err != nil {
		return nil, errno.ErrorVoConvertFailed("数据转换时出错").WithCause(err)
	}
	return reply, nil
}

func (svc *LoginService) RefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (*v1.RefreshTokenReply, error) {
	user := &biz.User{}
	if err := gconv.Struct(req, user); err != nil {
		return nil, errno.ErrorDoConvertFailed("数据转换时出错").WithCause(err)
	}
	data, err := svc.uc.RefreshToken(ctx)
	if err != nil {
		return nil, err
	}
	reply := &v1.RefreshTokenReply{}
	if err := gconv.Struct(data, &reply.Data); err != nil {
		return nil, errno.ErrorVoConvertFailed("数据转换时出错").WithCause(err)
	}
	return reply, nil
}

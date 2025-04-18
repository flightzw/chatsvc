// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.1
// source: chatsvc/v1/login.api.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLoginServiceLogin = "/chatsvc.v1.LoginService/Login"
const OperationLoginServiceRefreshToken = "/chatsvc.v1.LoginService/RefreshToken"
const OperationLoginServiceRegister = "/chatsvc.v1.LoginService/Register"

type LoginServiceHTTPServer interface {
	// Login 用户登录
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// RefreshToken 登录令牌续期
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenReply, error)
	// Register 用户注册
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
}

func RegisterLoginServiceHTTPServer(s *http.Server, srv LoginServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/chatsvc/v1/register", _LoginService_Register0_HTTP_Handler(srv))
	r.POST("/chatsvc/v1/login", _LoginService_Login0_HTTP_Handler(srv))
	r.GET("/chatsvc/v1/refresh-token", _LoginService_RefreshToken0_HTTP_Handler(srv))
}

func _LoginService_Register0_HTTP_Handler(srv LoginServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginServiceRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _LoginService_Login0_HTTP_Handler(srv LoginServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginServiceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _LoginService_RefreshToken0_HTTP_Handler(srv LoginServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefreshTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginServiceRefreshToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RefreshToken(ctx, req.(*RefreshTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefreshTokenReply)
		return ctx.Result(200, reply)
	}
}

type LoginServiceHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	RefreshToken(ctx context.Context, req *RefreshTokenRequest, opts ...http.CallOption) (rsp *RefreshTokenReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
}

type LoginServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewLoginServiceHTTPClient(client *http.Client) LoginServiceHTTPClient {
	return &LoginServiceHTTPClientImpl{client}
}

func (c *LoginServiceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/chatsvc/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLoginServiceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginServiceHTTPClientImpl) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...http.CallOption) (*RefreshTokenReply, error) {
	var out RefreshTokenReply
	pattern := "/chatsvc/v1/refresh-token"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLoginServiceRefreshToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginServiceHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/chatsvc/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLoginServiceRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

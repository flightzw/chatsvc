package service

import (
	"context"

	"github.com/flightzw/chatsvc/api/chatsvc/errno"
	v1 "github.com/flightzw/chatsvc/api/chatsvc/v1"
	"github.com/flightzw/chatsvc/internal/biz"
	"github.com/gogf/gf/v2/util/gconv"
)

type PrivateMessageService struct {
	v1.UnimplementedPrivateMessageServiceServer

	uc *biz.PrivateMessageUsecase
}

func NewPrivateMessageService(uc *biz.PrivateMessageUsecase) *PrivateMessageService {
	go uc.ProcessMessageSendResult(context.Background())
	return &PrivateMessageService{
		uc: uc,
	}
}

func (svc *PrivateMessageService) SendPrivateMessage(ctx context.Context, req *v1.SendPrivateMessageRequest) (*v1.SendPrivateMessageReply, error) {
	message := &biz.PrivateMessage{}
	if err := gconv.Struct(req, message); err != nil {
		return nil, errno.ErrorDoConvertFailed("数据转换时出错").WithCause(err)
	}
	data, err := svc.uc.SendPrivateMessage(ctx, message)
	if err != nil {
		return nil, err
	}
	reply := &v1.SendPrivateMessageReply{}
	if err := gconv.Struct(data, &reply.Data); err != nil {
		return nil, errno.ErrorVoConvertFailed("数据转换时出错").WithCause(err)
	}
	return reply, nil
}

func (svc *PrivateMessageService) RecallPrivateMessage(ctx context.Context, req *v1.RecallPrivateMessageRequest) (*v1.RecallPrivateMessageReply, error) {
	if err := svc.uc.RecallPrivateMessage(ctx, req.Id); err != nil {
		return nil, err
	}
	return &v1.RecallPrivateMessageReply{}, nil
}

func (svc *PrivateMessageService) PullOfflinePrivateMessage(ctx context.Context, req *v1.PullOfflineMessageRequest) (*v1.PullOfflineMessageReply, error) {
	if err := svc.uc.PullOfflinePrivateMessage(ctx, req.StartId); err != nil {
		return nil, err
	}
	return &v1.PullOfflineMessageReply{}, nil
}

func (svc *PrivateMessageService) ReadedPrivateMessage(ctx context.Context, req *v1.ReadedPrivateMessageRequest) (*v1.ReadedPrivateMessageReply, error) {
	if err := svc.uc.ReadedPrivateMessage(ctx, req.FriendId); err != nil {
		return nil, err
	}
	return &v1.ReadedPrivateMessageReply{}, nil
}

func (svc *PrivateMessageService) ListPrivateMessage(ctx context.Context, req *v1.ListPrivateMessageRequest) (*v1.ListPrivateMessageReply, error) {
	params := &biz.PrivateMessageQuery{}
	if err := gconv.Struct(req, params); err != nil {
		return nil, errno.ErrorDoConvertFailed("查询参数解析时出错").WithCause(err)
	}
	data, total, err := svc.uc.ListPrivateMessage(ctx, params, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListPrivateMessageReply{Total: total}
	if err = gconv.Struct(data, &reply.Data); err != nil {
		return nil, errno.ErrorVoConvertFailed("数据转换时出错").WithCause(err)
	}
	return reply, nil
}

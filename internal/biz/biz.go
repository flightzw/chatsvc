package biz

import (
	"context"

	"github.com/google/wire"

	"github.com/flightzw/chatsvc/internal/ws/client"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUsecase, NewFriendUsecase, NewPrivateMessageUsecase,
	client.InitChatClient, InitSensitiveWordFiliter)

// 事务
type TransactionInterface interface {
	Transaction(ctx context.Context, execute func(ctx context.Context) error) error // 启用事务
}

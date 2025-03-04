package ws

import (
	"github.com/flightzw/chatsvc/internal/enum"
)

// 待发送消息包装器
type MessageWrapper struct {
	RecvIds      []int32      `json:"recv_ids"`                // 消息接收者 ids
	Data         *SendMessage `json:"data"`                    // 消息数据
	NotifyResult bool         `json:"notify_result,omitempty"` // 是否通知发送结果
}

// 待发送消息
type SendMessage struct {
	Action enum.ActionType `json:"action"`
	Data   interface{}     `json:"data,omitempty"`
}

// 消息发送结果
type SendResult struct {
	Success bool         `json:"success"`
	Data    *SendMessage `json:"data,omitempty"`
}

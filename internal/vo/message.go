package vo

import (
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/flightzw/chatsvc/internal/enum"
)

type PrivateMessageVO struct {
	ID        int32              `json:"id"`
	SendID    int32              `json:"send_id"`    // 发送者 uid
	RecvID    int32              `json:"recv_id"`    // 接收者 uid
	Content   string             `json:"content"`    // 发送内容
	Type      enum.MessageType   `json:"type"`       // 消息类型
	Status    enum.MessageStatus `json:"status"`     // 状态 0:未送达 1:已送达 2:撤回 3:已读
	CreatedAt *gtime.Time        `json:"created_at"` // 创建时间
}

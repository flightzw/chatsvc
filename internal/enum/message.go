package enum

import "fmt"

// 通信操作类型
type ActionType int

const (
	ActionTypeAuth           = iota + 1 // 登录认证
	ActionTypeSignout                   // 注销
	ActionTypeOfflinePush               // 离线消息推送标记
	ActionTypePrivateMessage            // 私聊消息
	ActionTypeSystemMessage             // 系统消息
	ActionTypeMessageReaded             // 消息已读
)

var (
	actionTypeMap = map[ActionType]string{
		ActionTypeAuth:           "登录认证",
		ActionTypeSignout:        "注销",
		ActionTypeOfflinePush:    "消息推送通知",
		ActionTypePrivateMessage: "私聊消息",
		ActionTypeSystemMessage:  "系统消息",
		ActionTypeMessageReaded:  "消息已读",
	}
)

func (x ActionType) Map() string {
	if desc, ok := actionTypeMap[x]; ok {
		return fmt.Sprintf("%s(%d)", desc, x)
	}
	return fmt.Sprintf("UNKNOWN(%d)", x)
}

// 消息类型
type MessageType int

const (
	MessageTypeText   = iota + 1 // 文本消息
	MessageTypeImage             // 图片消息
	MessageTypeFile              // 文件消息
	MessageTypeVoice             // 语音消息
	MessageTypeVideo             // 视频消息
	MessageTypeSystem            // 系统消息
)

// 消息状态
type MessageStatus int

const (
	MessageStatusUnsend = iota // 未送达
	MessageStatusUnread        // 已送达-未读
	MessageStatusRecall        // 撤回
	MessageStatusReaded        // 已送达-已读
)

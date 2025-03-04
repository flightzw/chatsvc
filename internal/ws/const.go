package ws

const (
	RedisKeyServerIdCount     = "chatsvc:server-id:count"         // 聊天服务 id 计数
	RedisKeyUserServer        = "chatsvc:uid:%v:server"           // userSession/serverId 关联
	RedisKeyMessageQueue      = "chatsvc:server:%v:message-queue" // 消息投递队列
	RedisKeyResultNotifyQueue = "chatsvc:result-notify-queue"     // 投递结果通知队列
)

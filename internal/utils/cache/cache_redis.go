package cache

import "github.com/redis/go-redis/v9"

const (
	RedisKeyTodayNewUserCount    = "chatsvc:%s:new-users"           // 今日新增用户数
	RedisKeyTodayNewMessageCount = "chatsvc:%s:new-msgs"            // 今日新增消息数
	RedisKeyUserUpdatePassword   = "chatsvc:uid:%v:update-password" // 用户更新密码时间戳

	RedisKeyServerIdCount     = "chatsvc:server-id:count"               // 聊天服务 id 计数
	RedisKeyUserServer        = "chatsvc:uid:%v:server"                 // userSession/serverId 关联
	RedisKeyMessageQueue      = "chatsvc:server:%v:message-queue"       // 消息投递队列
	RedisKeyResultNotifyQueue = "chatsvc:result-notify-queue"           // 投递结果通知队列
	RedisKeyForceSignoutQueue = "chatsvc:server:%v:force-signout-queue" // 强制登出通知队列
)

// 新增计数并设置有效期
var IncrEX = redis.NewScript(`
local key = KEYS[1]
local count = redis.call('INCR', key)
if count == 1 then
    redis.call('EXPIRE', key, ARGV[1])
end
return count
`)

// 比较 keyval 并删除
var DelByValue = redis.NewScript(`
if redis.call('GET', KEYS[1]) == ARGV[1] then
    return redis.call('DEL', KEYS[1])
else
	return 0
end
`)

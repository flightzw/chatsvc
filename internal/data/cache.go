package data

import (
	"github.com/redis/go-redis/v9"
)

const (
	cacheKeyTodayNewUserCount    = "chatsvc:%s:new-users"
	cacheKeyTodayNewMessageCount = "chatsvc:%s:new-msgs"
)

// 新增计数并设置有效期
var incrEX = redis.NewScript(`
local key = KEYS[1]
local count = redis.call('INCR', key)
if count == 1 then
    redis.call('EXPIRE', key, ARGV[1])
end
return count
`)

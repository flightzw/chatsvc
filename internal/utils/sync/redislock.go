package sync

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	redis "github.com/redis/go-redis/v9"
)

const (
	_unlockScriptCmd = `if redis.call('get', KEYS[1]) == ARGV[1] 
	then 
		return redis.call('del', KEYS[1])
	else 
		return -1
	end
	`

	_expireScriptCmd = `if redis.call('get', KEYS[1]) == ARGV[1] 
	then 
		return redis.call('pexpire', KEYS[1], ARGV[2])
	else 
		return -1
	end
	`
)

type RedisLock struct {
	client     *redis.Client
	key, value string
	expire     time.Duration
	unlockSign chan struct{}
}

func NewRedisLock(client *redis.Client, key, value string, expire time.Duration) *RedisLock {
	return &RedisLock{
		client:     client,
		key:        key,
		value:      value,
		expire:     expire,
		unlockSign: make(chan struct{}),
	}
}

func (r *RedisLock) Lock(ctx context.Context) error {
	// 加锁
	res, err := r.client.SetNX(ctx, r.key, r.value, r.expire).Result()
	log.Println("r.client.SetNX res:", res, "err:", err)
	if err != nil {
		return errors.Wrap(err, "r.client.SetNX")
	}
	// 加锁成功后，创建守护线程，通过定时器给锁续期，直到收到锁释放信号
	go r.watchDog()
	return nil
}

func (r *RedisLock) UnLock() error {
	defer func() { r.unlockSign <- struct{}{} }()
	code, err := r.client.Eval(context.Background(), _unlockScriptCmd, []string{r.key}, r.value).Result()
	log.Println("r.client.Eval code:", code, "err:", err)
	if err != nil {
		return errors.Wrap(err, "r.client.Eval")
	}
	if code.(int64) == -1 {
		return errors.New("unlock failed")
	}
	return nil
}

func (r *RedisLock) watchDog() {
	for {
		select {
		case <-r.unlockSign:
			log.Println("lock unlocked, watchDog exit")
			return
		case <-time.After(r.expire / 2):
			code, err := r.client.Eval(context.Background(), _expireScriptCmd, []string{r.key}, r.value, r.expire/time.Millisecond).Result()
			if err != nil {
				log.Println("r.client.Expire err:", err)
			} else {
				log.Println("r.client.Expire finish, code:", code)
			}
		}
	}
}

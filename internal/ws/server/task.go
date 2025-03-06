package server

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/flightzw/chatsvc/internal/ws"
	"github.com/redis/go-redis/v9"
)

func listenMessagePushQueue(manager *MessageManager, gapTime time.Duration) {
	manager.log.Info("[message-manager] start listen message send queue.")
	queueKey := fmt.Sprintf(ws.RedisKeyMessageQueue, manager.serverID)
	ctx := context.Background()
	ticker := time.NewTicker(gapTime)
	defer ticker.Stop()

	for {
		<-ticker.C
		results, err := manager.redisClient.LPopCount(ctx, queueKey, 100).Result()
		if err != nil && err != redis.Nil {
			manager.log.Error("redisClient.LPopCount:", err)
			return
		}

		for _, result := range results {
			data := &ws.MessageWrapper{}
			if err = json.Unmarshal([]byte(result), data); err != nil {
				manager.log.Error("json.Unmarshal:", err)
			}
			manager.broadcast <- data
		}
	}
}

func listenForceSignoutQueue(hub *SessionHub, gapTime time.Duration) {
	hub.log.Info("[session-hub] start listen force signout queue.")
	queueKey := fmt.Sprintf(ws.RedisKeyForceSignoutQueue, hub.serverID)
	ctx := context.Background()
	ticker := time.NewTicker(gapTime)
	defer ticker.Stop()

	for {
		<-ticker.C
		results, err := hub.redisClient.LPopCount(ctx, queueKey, 100).Result()
		if err != nil && err != redis.Nil {
			hub.log.Error("redisClient.LPopCount:", err)
			return
		}

		for _, sessionID := range results {
			hub.removeSession(sessionID)
		}
	}
}

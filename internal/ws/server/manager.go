package server

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/flightzw/chatsvc/internal/ws"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type MessageManager struct {
	serverID    string
	log         *log.Helper
	redisClient *redis.Client
	broadcast   chan *ws.MessageWrapper
}

func NewMessageManager(logger log.Logger, redisClient *redis.Client) (*MessageManager, error) {
	serverID, err := redisClient.Incr(context.Background(), ws.RedisKeyServerIdCount).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redisClient.Incr")
	}
	manager := &MessageManager{
		log:         log.NewHelper(log.With(logger, "module", "chatsvc/ws/server")),
		serverID:    fmt.Sprint(serverID),
		redisClient: redisClient,
		broadcast:   make(chan *ws.MessageWrapper, 512),
	}
	// 定期从消息推送队列拉取数据
	go listenMessagePushQueue(manager, 100*time.Millisecond)

	return manager, nil
}

func (m *MessageManager) GetMessageSendChan(ctx context.Context) <-chan *ws.MessageWrapper {
	return m.broadcast
}

func (m *MessageManager) SendMessage(ctx context.Context, data *ws.MessageWrapper) error {
	keys := make([]string, 0, len(data.RecvIds))
	for _, recvId := range data.RecvIds {
		keys = append(keys, fmt.Sprintf(ws.RedisKeyUserServer, recvId))
	}
	serverIds, err := m.redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return errors.Wrap(err, "redisClient.MGet")
	}

	serverRecvMap := map[string][]int32{}
	for idx, recvId := range data.RecvIds {
		serverID, ok := serverIds[idx].(string)
		if !ok {
			m.log.Infof("user: %d not online, skip send step.", recvId)
			continue
		}
		serverRecvMap[serverID] = append(serverRecvMap[serverID], recvId)
	}

	for serverID, recvIds := range serverRecvMap {
		msg := &ws.MessageWrapper{
			RecvIds:      recvIds,
			Data:         data.Data,
			NotifyResult: data.NotifyResult,
		}
		msgBytes, _ := json.Marshal(msg)
		err = m.redisClient.RPush(ctx, fmt.Sprintf(ws.RedisKeyMessageQueue, serverID), string(msgBytes)).Err()
		if err != nil {
			return errors.Wrap(err, "redisClient.RPush")
		}
	}
	return nil
}

func (m *MessageManager) SendResultNotify(ctx context.Context, success bool, data *ws.SendMessage) error {
	resultBytes, _ := json.Marshal(ws.SendResult{
		Success: success,
		Data:    data,
	})
	if err := m.redisClient.RPush(ctx, ws.RedisKeyResultNotifyQueue, string(resultBytes)).Err(); err != nil {
		return errors.Wrap(err, "redisClient.RPush")
	}
	return nil
}

func (m *MessageManager) SendSignoutNotify(ctx context.Context, serverID, sessionID string) error {
	if err := m.redisClient.RPush(ctx, fmt.Sprintf(ws.RedisKeyForceSignoutQueue, serverID), sessionID).Err(); err != nil {
		return errors.Wrap(err, "redisClient.RPush")
	}
	return nil
}

func (c *MessageManager) IsOnline(ctx context.Context, sessionID string) (serverID string, ok bool) {
	serverID, err := c.redisClient.Get(ctx, fmt.Sprintf(ws.RedisKeyUserServer, sessionID)).Result()
	if err != nil {
		log.Error("redisClient.Get", err)
	}
	return serverID, err == nil
}
func (m *MessageManager) onSessionRegister(sessionId string) error {
	return m.redisClient.Set(context.Background(),
		fmt.Sprintf(ws.RedisKeyUserServer, sessionId), m.serverID, 3*time.Minute).Err()
}

func (m *MessageManager) onSessionHeartbeat(sessionId string, count int) error {
	if count%15 == 0 {
		return m.redisClient.Expire(context.Background(),
			fmt.Sprintf(ws.RedisKeyUserServer, sessionId), 3*time.Minute).Err()
	}
	return nil
}

func (m *MessageManager) onSessionLogout(sessionId string) error {
	return m.redisClient.Del(context.Background(), fmt.Sprintf(ws.RedisKeyUserServer, sessionId)).Err()
}

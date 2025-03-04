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
	serverId    string
	log         *log.Helper
	redisClient *redis.Client
	broadcast   chan *ws.MessageWrapper
}

func NewMessageManager(logger log.Logger, redisClient *redis.Client) (*MessageManager, error) {
	helper := log.NewHelper(log.With(logger, "module", "chatsvc/ws/server"))
	serverId, err := redisClient.Incr(context.Background(), ws.RedisKeyServerIdCount).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redisClient.Incr")
	}
	broadcast := make(chan *ws.MessageWrapper, 512)
	// 定期从消息推送队列拉取数据
	go func() {
		helper.Info("[message-manager] start listen message send queue.")
		queueKey := fmt.Sprintf(ws.RedisKeyMessageQueue, serverId)
		ctx := context.Background()
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			<-ticker.C
			results, err := redisClient.LPopCount(ctx, queueKey, 100).Result()
			if err != nil && err != redis.Nil {
				helper.Error("redisClient.LPopCount:", err)
				return
			}

			for _, result := range results {
				data := &ws.MessageWrapper{}
				if err = json.Unmarshal([]byte(result), data); err != nil {
					helper.Error("json.Unmarshal:", err)
				}
				broadcast <- data
			}
		}
	}()

	return &MessageManager{
		log:         helper,
		serverId:    fmt.Sprint(serverId),
		redisClient: redisClient,
		broadcast:   broadcast,
	}, nil
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
		serverId, ok := serverIds[idx].(string)
		if !ok {
			m.log.Infof("user: %d not online, skip send step.", recvId)
			continue
		}
		serverRecvMap[serverId] = append(serverRecvMap[serverId], recvId)
	}

	for serverId, recvIds := range serverRecvMap {
		msg := &ws.MessageWrapper{
			RecvIds:      recvIds,
			Data:         data.Data,
			NotifyResult: data.NotifyResult,
		}
		msgBytes, _ := json.Marshal(msg)
		err = m.redisClient.RPush(ctx, fmt.Sprintf(ws.RedisKeyMessageQueue, serverId), string(msgBytes)).Err()
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
	fmt.Println("[manager] send-result-notify:", string(resultBytes))
	if err := m.redisClient.RPush(ctx, ws.RedisKeyResultNotifyQueue, string(resultBytes)).Err(); err != nil {
		return errors.Wrap(err, "redisClient.RPush")
	}
	return nil
}

func (m *MessageManager) onSessionRegister(sessionId string) error {
	return m.redisClient.Set(context.Background(),
		fmt.Sprintf(ws.RedisKeyUserServer, sessionId), m.serverId, 3*time.Minute).Err()
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

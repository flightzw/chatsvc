package client

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

type ChatClient struct {
	log         *log.Helper
	redisClient *redis.Client
	notifyChan  chan *ws.SendResult
}

func InitChatClient(logger log.Logger, redisClient *redis.Client) *ChatClient {
	client := &ChatClient{
		log:         log.NewHelper(log.With(logger, "module", "chatsvc/ws/client")),
		redisClient: redisClient,
		notifyChan:  make(chan *ws.SendResult, 100),
	}
	// 定期从结果推送队列拉取数据
	go func() {
		client.log.Info("[chat-client] start listen message send result notify queue.")
		ctx := context.Background()
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			<-ticker.C
			results, err := redisClient.LPopCount(ctx, ws.RedisKeyResultNotifyQueue, 100).Result()
			if err != nil && err != redis.Nil {
				client.log.Error("redisClient.LPopCount:", err)
				return
			}

			for _, result := range results {
				fmt.Println("[client] result-notify:", result)
				data := &ws.SendResult{}
				if err = json.Unmarshal([]byte(result), &data); err != nil {
					client.log.Error("json.Unmarshal:", err)
					return
				}
				client.notifyChan <- data
			}
		}
	}()
	return client
}

func (c *ChatClient) GetResultNotifyChan(ctx context.Context) <-chan *ws.SendResult {
	return c.notifyChan
}

// 检查在线状态，返回在线 userIds
func (c *ChatClient) IsOnline(ctx context.Context, userIds ...int32) map[int32]bool {
	keys := make([]string, 0, len(userIds))
	for _, userID := range userIds {
		keys = append(keys, fmt.Sprintf(ws.RedisKeyUserServer, userID))
	}
	onlineMap := map[int32]bool{}
	results, err := c.redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		log.Error("redisClient.MGet", err)
		return onlineMap
	}
	for idx, userID := range userIds {
		if results[idx] != nil {
			onlineMap[userID] = true
		}
	}
	return onlineMap
}

// 推送消息数据
func (m *ChatClient) SendMessage(ctx context.Context, data *ws.MessageWrapper) error {
	keys := make([]string, 0, len(data.RecvIds))
	for _, recvID := range data.RecvIds {
		keys = append(keys, fmt.Sprintf(ws.RedisKeyUserServer, recvID))
	}
	serverIds, err := m.redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return errors.Wrap(err, "redisClient.MGet")
	}

	serverRecvMap := map[string][]int32{}
	for idx, recvID := range data.RecvIds {
		serverId, ok := serverIds[idx].(string)
		if !ok {
			m.log.Infof("user: %d not online, skip send step.", recvID)
			continue
		}
		serverRecvMap[serverId] = append(serverRecvMap[serverId], recvID)
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

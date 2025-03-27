package server

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"

	"github.com/flightzw/chatsvc/internal/ws"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 5 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type userSession struct {
	id         string
	createTime *gtime.Time
	hub        *SessionHub
	conn       *websocket.Conn
	aiConn     *websocket.Conn
	sendChan   chan *ws.MessageWrapper
}

// 接收并转发来自客户端的消息
func (c *userSession) sendMessage(ctx context.Context) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	count := 0
	connID := fmt.Sprintf("%s:%s:%d", c.hub.serverID, c.id, c.createTime.UnixNano())

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		count++
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err := c.hub.onSessionHeartbeat(c.id, count); err != nil {
			c.hub.log.Errorf("session [%s] heartbeat callback failed, error: %v", connID, err)
		}
		return nil
	})

	for {
		wrapper := &ws.MessageWrapper{}
		err := c.conn.ReadJSON(wrapper)
		if err != nil {
			c.hub.log.Errorf("session [%s] read message failed, error: %v", connID, err)
			break
		}
		if err = c.hub.SendMessage(ctx, wrapper); err != nil {
			c.hub.log.Warnf("[%s] SendMessage failed: %v", err)
		}
	}
}

// 将消息推送到客户端
func (c *userSession) recvMessage(ctx context.Context) {
	connID := fmt.Sprintf("%s:%s:%d", c.hub.serverID, c.id, c.createTime.UnixNano())
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case wrapper, ok := <-c.sendChan:
			// 从 chan 接收消息发送至客户端
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			wrappers := []*ws.MessageWrapper{wrapper}
			for range len(c.sendChan) {
				wrappers = append(wrappers, <-c.sendChan)
			}
			var (
				code   int32
				reason string
			)
			for _, wrapper := range wrappers {
				startTime := time.Now()

				err := c.conn.WriteJSON(wrapper.Data)
				if wrapper.NotifyResult {
					if resErr := c.hub.SendResultNotify(ctx, err == nil, wrapper.Data); resErr != nil {
						c.hub.log.Errorf("session [%s] send result notify failed, error: %v", connID, err)
					}
				}
				if se := errors.FromError(err); se != nil {
					code = se.Code
					reason = se.Reason
				}
				level, stack := extractError(err)
				c.hub.log.Log(level,
					"uid", c.id,
					"action", wrapper.Data.Action.Map(),
					"data", toJson(wrapper.Data.Data),
					"code", code,
					"reason", reason,
					"stack", stack,
					"latency", time.Since(startTime),
				)
			}
		case <-ticker.C:
			// 定时发送心跳
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}

func toJson(v interface{}) string {
	bts, _ := json.Marshal(v)
	return string(bts)
}

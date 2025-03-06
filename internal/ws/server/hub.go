package server

import (
	"context"
	"strconv"
	"time"

	"github.com/flightzw/chatsvc/internal/enum"
	"github.com/flightzw/chatsvc/internal/ws"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type SessionHub struct {
	*MessageManager

	sessions   map[string]*userSession // 本地会话 map
	register   chan *userSession       // 注册 chan
	unregister chan string             // 注销 chan
}

func InitSessionHub(logger log.Logger, redisClient *redis.Client) (*SessionHub, error) {
	manager, err := NewMessageManager(logger, redisClient)
	if err != nil {
		return nil, err
	}
	hub := &SessionHub{
		MessageManager: manager,
		sessions:       make(map[string]*userSession),
		register:       make(chan *userSession),
		unregister:     make(chan string),
	}
	// 监听本地消息推送队列，转发数据
	go hub.Run()
	// 监听强制登出信号队列，移除连接
	go listenForceSignoutQueue(hub, 100*time.Millisecond)
	return hub, nil
}

func (h *SessionHub) Run() {
	h.log.Info("[session-hub] start listen register/unregister/broadcast channel.")
	broadcast := h.GetMessageSendChan(context.Background())
	for {
		select {
		case session := <-h.register:
			h.sessions[session.id] = session
			if err := h.onSessionRegister(session.id); err != nil {
				h.log.Error("manager.onSessionRegister:", err)
			}
		case sessionId := <-h.unregister:
			session, ok := h.sessions[sessionId]
			if !ok {
				break
			}
			delete(h.sessions, sessionId)
			close(session.sendChan)
			if err := h.onSessionLogout(sessionId); err != nil {
				h.log.Error("manager.onSessionLogout:", err)
			}
		case wrapper := <-broadcast:
			for _, recvId := range wrapper.RecvIds {
				sessionId := strconv.Itoa(int(recvId))
				if session, ok := h.sessions[sessionId]; ok {
					select {
					case session.sendChan <- wrapper:
					default:
						close(session.sendChan)
						delete(h.sessions, sessionId)
					}
				}
			}
		}
	}
}

func (hub *SessionHub) removeSession(sessionID string) {
	session := hub.sessions[sessionID]
	if session == nil {
		return
	}
	err := session.conn.WriteJSON(ws.SendMessage{
		Action: enum.ActionTypeSignout,
		Data:   "已在其他地方登录，将强制退出",
	})
	hub.log.Infof("已通知用户断开连接 [%s:%s:%d], res: %v", hub.serverID, sessionID, session.createTime.UnixNano(), err)
	session.conn.Close()
}

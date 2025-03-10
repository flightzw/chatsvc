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
	unregister chan *userSession       // 注销 chan
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
		unregister:     make(chan *userSession),
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
			err := h.onSessionRegister(session.id)
			h.log.Infof("[session-hub] session [%s:%s:%d] register finished, error: %v", h.serverID, session.id, session.createTime.UnixNano(), err)
		case session := <-h.unregister:
			storeSession, ok := h.sessions[session.id]
			if !ok || storeSession != session {
				h.log.Infof("[session-hub] session [%s:%s:%d] has been unregister, continue", h.serverID, session.id, session.createTime.UnixNano())
				break
			}
			delete(h.sessions, session.id)
			close(session.sendChan)
			err := h.onSessionLogout(session.id)
			h.log.Infof("[session-hub] session [%s:%s:%d] unregister finished, error: %v", h.serverID, session.id, session.createTime.UnixNano(), err)
		case wrapper := <-broadcast:
			for _, recvId := range wrapper.RecvIds {
				sessionID := strconv.Itoa(int(recvId))
				if session, ok := h.sessions[sessionID]; ok {
					select {
					case session.sendChan <- wrapper:
					default:
						close(session.sendChan)
						delete(h.sessions, sessionID)
					}
				}
			}
		}
	}
}

func (hub *SessionHub) removeSession(sessionID string) {
	session := hub.sessions[sessionID]
	if session == nil {
		hub.log.Infof("session [%s:%s] not found, no need signout", hub.serverID, sessionID)
		return
	}
	err := session.conn.WriteJSON(ws.SendMessage{
		Action: enum.ActionTypeSignout,
		Data:   "已在其他地方登录，将强制退出",
	})
	hub.log.Infof("session [%s:%s:%d] signout message send finished, error: %v", hub.serverID, sessionID, session.createTime.UnixNano(), err)
	session.conn.Close()
}

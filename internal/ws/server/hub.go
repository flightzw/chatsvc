package server

import (
	"context"
	"strconv"

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

// 核心
// 消息推送、消息接收模块
// 消息处理函数
/*
type Sender interface {
	// inner queue, send message to queue
	func(ctx context.Context, data *sendInfo) error
}

type Receiver interface {
	// inner queue, receive message from queue
	func(ctx context.Context, func()) error
}
流程：
发送消息到队列 Sender, Receiver
监听队列状态，接收消息并分发到 client 局部队列 func(clientMap [string]*client)
client 监听局部队列转发消息到客户端
*/

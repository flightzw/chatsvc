package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	authjwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"

	"github.com/flightzw/chatsvc/internal/entity"
	"github.com/flightzw/chatsvc/internal/enum"
	"github.com/flightzw/chatsvc/internal/ws"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type AuthFunc func(token string) (id string, claims jwt.Claims, err error)

func MakeChatHandleFunc(hub *SessionHub, authFunc AuthFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		conn, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			hub.log.Error("websocket upgrade failed:", err)
			return
		}

		loginChan := make(chan ws.SendMessage)
		defer close(loginChan)
		go func(result chan<- ws.SendMessage) {
			msg := ws.SendMessage{}
			if err := conn.ReadJSON(&msg); err != nil {
				hub.log.Error("read auth message failed:", err)
				return
			}
			result <- msg
		}(loginChan)

		var (
			sessionID  string
			claims     jwt.Claims
			createTime *gtime.Time
		)
		select {
		case <-time.NewTimer(pongWait).C:
			sendErrorMessage(conn, "身份认证等待超时")
			conn.Close()
			return
		case msg := <-loginChan:
			if msg.Action != enum.ActionTypeAuth {
				sendErrorMessage(conn, "尚未完成身份认证，操作无效")
				conn.Close()
				return
			}
			token, _ := msg.Data.(map[string]any)["token"].(string)
			sessionID, claims, err = authFunc(token)
			if err != nil {
				code, resMsg := 0, ""
				if errors.Is(err, jwt.ErrTokenExpired) {
					code, resMsg = 401, "token 已过期"
				} else {
					code, resMsg = 400, "token 无效"
				}
				err1 := conn.WriteJSON(ws.SendMessage{
					Action: enum.ActionTypeAuth,
					Data:   entity.AnyMap{"code": code, "message": resMsg},
				})
				hub.log.Errorf("auth failed: %v, resp error: %v", err, err1)
				return
			}
		}
		serverID, ok := hub.IsOnline(context.Background(), sessionID)
		if serverID == hub.serverID {
			// 存在本地会话，直接注销
			hub.removeSession(sessionID)
		} else if ok {
			// 存在远程会话，发送通知
			if err = hub.SendSignoutNotify(context.Background(), serverID, sessionID); err != nil {
				hub.log.Errorf("session: [%s:%s] send signout notify failed, error: %v", serverID, sessionID, err)
				conn.Close()
				return
			}
		}
		if err = conn.WriteJSON(ws.SendMessage{Action: enum.ActionTypeAuth}); err != nil {
			return
		}
		createTime = gtime.Now()
		hub.log.Infof("session [%s:%s:%d] auth passed.", hub.serverID, sessionID, createTime.UnixNano())

		session := &userSession{
			id:         sessionID,
			hub:        hub,
			conn:       conn,
			sendChan:   make(chan *ws.MessageWrapper, 256),
			createTime: createTime,
		}
		hub.register <- session

		ctx := authjwt.NewContext(context.Background(), claims)
		go session.sendMessage(ctx)
		go session.recvMessage(ctx)
	}
}

func sendErrorMessage(conn *websocket.Conn, msg string) {
	conn.WriteJSON(ws.SendMessage{
		Action: enum.ActionTypeSystemMessage,
		Data:   msg,
	})
}

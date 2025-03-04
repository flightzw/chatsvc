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
			sessionId  string
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
			code, resMsg := 0, "认证通过"
			token, _ := msg.Data.(map[string]any)["token"].(string)
			sessionId, claims, err = authFunc(token)
			if err != nil {
				hub.log.Errorf("auth failed:", err)
				if errors.Is(err, jwt.ErrTokenExpired) {
					code, resMsg = 401, "token 已过期"
				} else {
					code, resMsg = 400, "token 无效"
				}
				conn.WriteJSON(ws.SendMessage{
					Action: enum.ActionTypeAuth,
					Data:   entity.AnyMap{"code": code, "message": resMsg},
				})
				return
			}

			if oldClient := hub.sessions[sessionId]; oldClient != nil {
				oldClient.conn.WriteJSON(ws.SendMessage{
					Action: enum.ActionTypeSignout,
					Data:   "已建立新连接，断开通信",
				})
			}
			if err = conn.WriteJSON(ws.SendMessage{Action: enum.ActionTypeAuth}); err != nil {
				return
			}
			createTime = gtime.Now()
			hub.log.Infof("[%s:%d] 身份认证已通过", sessionId, createTime.UnixNano())
		}

		session := &userSession{
			id:         sessionId,
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

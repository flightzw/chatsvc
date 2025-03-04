package server

import (
	"errors"
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	authjwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"

	v1 "github.com/flightzw/chatsvc/api/chatsvc/v1"
	"github.com/flightzw/chatsvc/internal/conf"
	"github.com/flightzw/chatsvc/internal/middleware"
	"github.com/flightzw/chatsvc/internal/service"
	"github.com/flightzw/chatsvc/internal/utils/jwt"
	"github.com/flightzw/chatsvc/internal/ws/server"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger,
	hub *server.SessionHub,
	friendSvc *service.FriendService,
	userSvc *service.UserService,
	loginSvc *service.LoginService,
	messageSvc *service.PrivateMessageService,
) *http.Server {
	signMethod := jwtv5.SigningMethodRS256
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.Server(log.NewFilter(logger, log.FilterFunc(makeLogFilter()))),
			middleware.MakeJwtMiddleware(signMethod, c.Jwt.AccessToken.Pubfile, jwt.NewClaims, []string{
				v1.OperationLoginServiceLogin,
				v1.OperationLoginServiceRegister,
				v1.OperationLoginServiceRefreshToken,
			}, false),
			middleware.MakeJwtMiddleware(signMethod, c.Jwt.RefreshToken.Pubfile, jwt.NewClaims, []string{
				v1.OperationLoginServiceRefreshToken,
			}, true),
			middleware.Validator(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"PUT", "POST", "GET", "DELETE", "OPTIONS"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	v1.RegisterUserServiceHTTPServer(srv, userSvc)
	v1.RegisterFriendServiceHTTPServer(srv, friendSvc)
	v1.RegisterLoginServiceHTTPServer(srv, loginSvc)
	v1.RegisterPrivateMessageServiceHTTPServer(srv, messageSvc)

	srv.HandleFunc("/ping", func(w nethttp.ResponseWriter, r *nethttp.Request) { w.Write([]byte("pong")) })
	srv.HandleFunc("/chatsvc/chat/channel", server.MakeChatHandleFunc(hub,
		makeChatAuthFunc(signMethod, middleware.GetKeyFunc(signMethod, c.Jwt.AccessToken.Pubfile))))
	return srv
}

func makeChatAuthFunc(signMethod jwtv5.SigningMethod, keyFunc jwtv5.Keyfunc) server.AuthFunc {
	return func(token string) (id string, claims jwtv5.Claims, err error) {
		claimsImpl := &jwtv5.RegisteredClaims{}
		tokenInfo, err := jwtv5.ParseWithClaims(token, claimsImpl, keyFunc)
		if err != nil {
			if errors.Is(err, jwtv5.ErrTokenMalformed) || errors.Is(err, jwtv5.ErrTokenUnverifiable) {
				return "", nil, authjwt.ErrTokenInvalid
			}
			if errors.Is(err, jwtv5.ErrTokenNotValidYet) || errors.Is(err, jwtv5.ErrTokenExpired) {
				return "", nil, authjwt.ErrTokenExpired
			}
			return "", nil, authjwt.ErrTokenParseFail
		}
		if !tokenInfo.Valid {
			return "", nil, authjwt.ErrTokenInvalid
		}
		if tokenInfo.Method != signMethod {
			return "", nil, authjwt.ErrUnSupportSigningMethod
		}
		return claimsImpl.ID, claimsImpl, nil
	}
}

func serveHome(w http.ResponseWriter, r *nethttp.Request) {
	if r.URL.Path != "/" {
		nethttp.Error(w, "Not found", nethttp.StatusNotFound)
		return
	}
	if r.Method != nethttp.MethodGet {
		nethttp.Error(w, "Method not allowed", nethttp.StatusMethodNotAllowed)
		return
	}
	nethttp.ServeFile(w, r, "./web/chat-client/home.html")
}

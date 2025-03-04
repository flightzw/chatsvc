package middleware

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// 配置 Jwt 认证中间件
//
//	secret    加密密钥
//	method    加密算法
//	newClaims WithClaims 参数
//	paths     接口 urls
//	reverse   反转标识: 为 true 表示仅为 paths 中的接口配置认证中间件
func MakeJwtMiddleware(method jwtv5.SigningMethod, pubfile string, newClaims func() jwtv5.Claims, paths []string, reverse bool) middleware.Middleware {
	return selector.Server(
		jwt.Server(
			GetKeyFunc(method, pubfile),
			jwt.WithSigningMethod(method),
			jwt.WithClaims(newClaims),
		),
	).Match(makeAuthWhiteListMatchFunc(paths, reverse)).Build()
}

func makeAuthWhiteListMatchFunc(paths []string, reverse bool) selector.MatchFunc {
	nilStruct := struct{}{}
	pathMap := make(map[string]struct{})
	for _, path := range paths {
		pathMap[path] = nilStruct
	}
	return func(ctx context.Context, operation string) bool {
		if _, ok := pathMap[operation]; ok {
			return reverse
		}
		return !reverse
	}
}

func GetKeyFunc(method jwtv5.SigningMethod, pubfile string) jwtv5.Keyfunc {
	var (
		err       error
		publicKey any
	)
	content, _ := os.ReadFile(pubfile)
	switch method {
	case jwtv5.SigningMethodRS256:
		publicKey, err = jwtv5.ParseRSAPublicKeyFromPEM(content)
	}
	return func(*jwtv5.Token) (any, error) {
		return publicKey, err
	}
}

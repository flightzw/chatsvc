package middleware

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/bufbuild/protovalidate-go"
	"github.com/go-kratos/kratos/v2/middleware"
	authjwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"

	"github.com/flightzw/chatsvc/api/chatsvc/errno"
	"github.com/flightzw/chatsvc/internal/utils/cache"
	"github.com/flightzw/chatsvc/internal/utils/jwt"
)

// Validator is a validator middleware.
func Validator(client *redis.Client) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if claims, _ := jwt.GetRegisteredClaims(ctx); claims != nil {
				timestamp, err := client.Get(ctx, fmt.Sprintf(cache.RedisKeyUserUpdatePassword, claims.ID)).Result()
				if err != nil && !errors.Is(err, redis.Nil) {
					return nil, errno.ErrorSystemInternalFailure("此功能暂不可用").WithCause(err)
				}
				if timestamp != "" && timestamp > strconv.FormatInt(claims.IssuedAt.Unix(), 10) {
					return nil, authjwt.ErrTokenInvalid
				}
			}

			if err = protovalidate.Validate(req.(proto.Message)); err != nil {
				return nil, errno.ErrorParamInvalid("参数校验时出错").WithMetadata(generateErrorMetadata(err.Error()))
			}
			return handler(ctx, req)
		}
	}
}

func generateErrorMetadata(message string) map[string]string {
	dataMap := map[string]string{}
	for _, row := range strings.Split(message, "\n")[1:] {
		items := strings.Split(row, " ")
		key := items[2][:len(items[2])-1]
		if _, ok := dataMap[key]; ok {
			dataMap[key] += ","
		}
		dataMap[key] += items[3]
	}
	return dataMap
}

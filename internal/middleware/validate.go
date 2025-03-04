package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/bufbuild/protovalidate-go"
	"github.com/flightzw/chatsvc/api/chatsvc/errno"
	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/protobuf/proto"
)

// Validator is a validator middleware.
func Validator() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if err = protovalidate.Validate(req.(proto.Message)); err != nil {
				return nil, errno.ErrorParamInvalid("参数校验时出错").WithMetadata(generateErrorMetadata(err.Error()))
			}
			return handler(ctx, req)
		}
	}
}

func generateErrorMetadata(message string) map[string]string {
	fmt.Println("generateErrorMetadata:", message)
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

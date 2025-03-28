package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http/status"
	"google.golang.org/grpc/codes"
)

type logInfoKey struct{}

func Logger(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			var (
				startTime = time.Now()
				code      = int32(status.FromGRPCCode(codes.OK))
				infoMap   = map[string]any{}
				reason    string
				kind      string
				operation string
			)
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(context.WithValue(ctx, logInfoKey{}, infoMap), req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			level, stack := extractError(err)
			log.NewHelper(log.WithContext(ctx, logger)).Log(level,
				"kind", "server",
				"component", kind,
				"operation", operation,
				"uid", infoMap["uid"],
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"stack", stack,
				"latency", time.Since(startTime),
			)
			return
		}
	}
}

func extractArgs(req any) string {
	args, _ := json.Marshal(req)
	return string(args)
}

func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}

func SetLogInfoToContext(ctx context.Context, userID int32) {
	infoMap, ok := ctx.Value(logInfoKey{}).(map[string]any)
	if ok {
		infoMap["uid"] = userID
	}
}

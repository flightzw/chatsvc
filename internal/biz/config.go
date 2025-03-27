package biz

import "context"

type ConfigRepo interface {
	GetConfigMap(ctx context.Context, parentID int32) (data map[string]string, err error)
}

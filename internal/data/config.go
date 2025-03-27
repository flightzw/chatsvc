package data

import (
	"context"
	"fmt"

	"github.com/flightzw/chatsvc/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/pkg/errors"
)

type configRepo struct {
	*Data

	log *log.Helper
}

// NewConfigRepo .
func NewConfigRepo(data *Data, logger log.Logger) biz.ConfigRepo {
	return &configRepo{
		Data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *configRepo) GetConfigMap(ctx context.Context, parentID int32) (data map[string]string, err error) {
	var (
		cacheKey = fmt.Sprintf("config:group_id:%d", parentID)
		config   = repo.UseQuery(ctx).Config
		configDo = config.WithContext(ctx)
	)
	if value, ok := repo.cache.Get(cacheKey); ok {
		return value.(map[string]string), nil
	}

	result, err := configDo.Where(config.ParentID.Eq(parentID)).Find()
	if err != nil {
		return nil, errors.Wrap(err, "configDo.Find")
	}
	data = map[string]string{}
	for _, res := range result {
		data[res.Key] = res.Value
	}
	repo.cache.Set(cacheKey, data, 0)
	return data, nil
}

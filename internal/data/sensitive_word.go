package data

import (
	"context"

	"github.com/flightzw/chatsvc/internal/biz"
	"github.com/flightzw/chatsvc/internal/biz/query"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"gorm.io/gen"
)

type sensitiveWordRepo struct {
	*Data
	log *log.Helper
}

func NewSensitiveWordRepo(data *Data, logger log.Logger) biz.SensitiveWordRepo {
	return &sensitiveWordRepo{
		Data: data,
		log:  log.NewHelper(log.With(logger, "module", "chatsvc/data/SensitiveWordRepo")),
	}
}

func (repo *sensitiveWordRepo) ListSensitiveWord(ctx context.Context, queryFunc query.QueryFunc, page, pageSize int) (data []*biz.SensitiveWord, total int64, err error) {
	sw := repo.UseQuery(ctx).SensitiveWord
	swDo := sw.WithContext(ctx)
	data = []*biz.SensitiveWord{}

	swDo.DO = *queryFunc(&swDo.DO).(*gen.DO)
	result, total, err := swDo.FindByPage((page-1)*pageSize, pageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "swDo.FindByPage")
	}
	if err = gconv.Structs(result, &data); err != nil {
		return nil, 0, errors.Wrap(err, "gconv.Structs")
	}
	return data, total, nil
}

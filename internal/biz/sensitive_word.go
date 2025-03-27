package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/pkg/errors"

	"github.com/flightzw/chatsvc/internal/biz/query"
	"github.com/flightzw/chatsvc/internal/utils/stringx"
)

type SensitiveWord struct {
	ID        int32       `json:"id"`
	Content   string      `json:"content"`    // 敏感词内容
	Enabled   bool        `json:"enabled"`    // 是否启用 1:是 0:否
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 更新时间
}

type SensitiveWordRepo interface {
	ListSensitiveWord(ctx context.Context, queryFunc query.QueryFunc, page, pageSize int) ([]*SensitiveWord, int64, error)
}

func InitSensitiveWordFiliter(logger log.Logger, repo SensitiveWordRepo) (*stringx.Filter, error) {
	var (
		ctx            = context.Background()
		filter         = stringx.NewFilter()
		sw             = query.NewSensitiveWordQuery()
		page, pageSize = 1, 1000
		helper         = log.NewHelper(log.With(logger, "func", "chatsvc/biz/InitSensitiveWordFiliter"))
	)
	data, _, err := repo.ListSensitiveWord(ctx, func(do query.QueryChain) query.QueryChain { return do.Where(sw.Enabled.Eq(1)) }, page, pageSize)
	if err != nil {
		return nil, errors.Wrap(err, "init sensitive word filiter failed")
	}
	words := make([]string, 0, len(data))
	for _, word := range data {
		words = append(words, word.Content)
	}
	filter.AddWord(words...)

	go func() {
		var (
			data   []*SensitiveWord
			err    error
			ticker = time.NewTicker(60 * time.Second)
		)
		defer ticker.Stop()

		helper.Info("start sync filter sensitive words config...")
		for {
			<-ticker.C
			qf := func(do query.QueryChain) query.QueryChain {
				return do.Where(sw.UpdatedAt.Gt(time.Now().Add(-1 * time.Minute)))
			}
			data, _, err = repo.ListSensitiveWord(ctx, qf, page, pageSize)
			if err != nil {
				helper.Error("sync filter sensitive words config failed:", err)
				continue
			}
			adds, dels := []string{}, []string{}
			for _, word := range data {
				if word.Enabled {
					adds = append(adds, word.Content)
				} else {
					dels = append(dels, word.Content)
				}
			}
			addNum, delNum := len(adds), len(dels)
			if addNum > 0 {
				filter.AddWord(adds...)
			}
			if delNum > 0 {
				filter.DelWord(dels...)
			}
			if addNum > 0 || delNum > 0 {
				helper.Infof("sync filter sensitive words config success, add: %d, del: %d", len(adds), len(dels))
			}
		}
	}()
	return filter, nil
}

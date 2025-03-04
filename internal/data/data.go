package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/flightzw/chatsvc/internal/conf"
	"github.com/flightzw/chatsvc/internal/data/query"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedisClient, NewUserRepo, NewFriendRepo, NewPrivateMessageRepo)

type queryKey struct{}

// Data .
type Data struct {
	// TODO wrapped database client
	db    *gorm.DB
	query *query.Query
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "chatsvc/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	if conf.Env != "production" {
		db = db.Debug()
	}
	return db
}

func NewRedisClient(conf *conf.Data) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       int(conf.Redis.Db),
	})
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:    db,
		query: query.Use(db),
	}, cleanup, nil
}

func (t *Data) UseQuery(ctx context.Context) *query.Query {
	queryTx, ok := ctx.Value(queryKey{}).(*query.Query)
	if !ok {
		return t.query
	}
	return queryTx
}

func (t *Data) Transaction(ctx context.Context, execute func(ctx context.Context) error) error {
	return t.query.Transaction(func(tx *query.Query) error {
		return execute(context.WithValue(ctx, queryKey{}, tx))
	})
}

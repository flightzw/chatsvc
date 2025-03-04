package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"gorm.io/gen"

	"github.com/flightzw/chatsvc/internal/biz"
	"github.com/flightzw/chatsvc/internal/biz/query"
	"github.com/flightzw/chatsvc/internal/data/model"
	"github.com/flightzw/chatsvc/internal/entity"
)

type userRepo struct {
	*Data
	log *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		Data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *userRepo) CreateUser(ctx context.Context, user *biz.User) (id int32, err error) {
	userDo := repo.UseQuery(ctx).User.WithContext(ctx)
	data := &model.User{}

	if err = gconv.Struct(user, data); err != nil {
		return 0, errors.Wrap(err, "gconv.Struct")
	}
	if err = userDo.Create(data); err != nil {
		return 0, errors.Wrap(err, "userDo.Create")
	}
	return data.ID, nil
}

func (repo *userRepo) GetUser(ctx context.Context, id int32) (data *biz.User, err error) {
	user := repo.UseQuery(ctx).User
	userDo := user.WithContext(ctx)
	data = &biz.User{}

	userData, err := userDo.Where(user.ID.Eq(id)).First()
	if err != nil {
		return nil, errors.Wrap(err, "userDo.First")
	}
	if err = gconv.Struct(userData, data); err != nil {
		return nil, errors.Wrap(err, "gconv.Struct")
	}
	return data, nil
}

func (repo *userRepo) GetUserByUsername(ctx context.Context, username string) (data *biz.User, err error) {
	user := repo.UseQuery(ctx).User
	userDo := user.WithContext(ctx)
	data = &biz.User{}

	userData, err := userDo.Where(user.Username.Eq(username)).First()
	if err != nil {
		return nil, errors.Wrap(err, "userDo.First")
	}
	if err = gconv.Struct(userData, data); err != nil {
		return nil, errors.Wrap(err, "gconv.Struct")
	}
	return data, nil
}

func (repo *userRepo) ListUser(ctx context.Context, queryFunc query.QueryFunc, page, pageSize int) (data []*biz.User, total int64, err error) {
	user := repo.UseQuery(ctx).User
	userDo := user.WithContext(ctx)
	data = []*biz.User{}

	userDo.DO = *queryFunc(&userDo.DO).(*gen.DO)
	result, total, err := userDo.FindByPage((page-1)*pageSize, pageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "userDo.FindByPage")
	}
	if err = gconv.Structs(result, &data); err != nil {
		return nil, 0, errors.Wrap(err, "gconv.Structs")
	}
	return data, total, nil
}

func (repo *userRepo) UpdateUser(ctx context.Context, id int32, data entity.AnyMap) (err error) {
	user := repo.UseQuery(ctx).User
	userDo := user.WithContext(ctx).Omit(user.ID, user.CreatedAt, user.DeletedAt)

	if _, err = userDo.Where(user.ID.Eq(id)).Updates(data.Assert()); err != nil {
		return errors.Wrap(err, "userDo.Updates")
	}
	return nil
}

func (repo *userRepo) DeleteUser(ctx context.Context, id int32) (err error) {
	panic("not implemented") // TODO: Implement
}

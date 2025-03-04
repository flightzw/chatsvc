package query

import (
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type QueryChain interface {
	gen.SubQuery

	Not(conds ...gen.Condition) gen.Dao
	Or(conds ...gen.Condition) gen.Dao

	Select(columns ...field.Expr) gen.Dao
	Where(conds ...gen.Condition) gen.Dao
	Order(columns ...field.Expr) gen.Dao
}

type QueryFunc func(do QueryChain) QueryChain

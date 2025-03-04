package data

import (
	"gorm.io/gen"
)

type queryCondition[S any] interface {
	Where(conds ...gen.Condition) S
	Or(conds ...gen.Condition) S
}

func parseConditions[T queryCondition[T]](dataObj T, conds []any) T {
	if len(conds) == 0 {
		return dataObj
	}
	// first group conditions
	switch v := conds[0].(type) {
	case []gen.Condition:
		dataObj = dataObj.Where(v...)
	case gen.Condition:
		dataObj = dataObj.Where(v)
	default:
		panic("The conds's type not supported")
	}
	// other conditions
	for _, cond := range conds[1:] {
		switch v := cond.(type) {
		case []gen.Condition:
			dataObj = dataObj.Or(v...)
		case gen.Condition:
			dataObj = dataObj.Or(v)
		default:
			panic("The conds's type not supported")
		}
	}
	return dataObj
}

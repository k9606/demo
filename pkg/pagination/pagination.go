package pagination

import (
	"gorm.io/gorm"
)

type Param struct {
	Page    int
	Limit   int
	Count   int64
	Keyword string
}

func Paginate(db *gorm.DB, param Param) (*gorm.DB, int64) {
	var (
		page  = param.Page
		limit = param.Limit
		count = param.Count
	)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 100 {
		limit = 1
	}

	db.Count(&count).Offset((page - 1) * limit).Limit(limit)

	return db, count
}

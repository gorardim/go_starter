package g

import "gorm.io/gorm"

var _ Builder = (*LimitBuilder)(nil)

type LimitBuilder struct {
	Limit int
}

func Limit(limit int) LimitBuilder {
	return LimitBuilder{
		Limit: limit,
	}
}

func (l LimitBuilder) Build(db *gorm.DB) *gorm.DB {
	return db.Limit(l.Limit)
}

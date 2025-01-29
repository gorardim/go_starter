package g

import "gorm.io/gorm"

var _ Builder = (*WhereBuilder)(nil)

type WhereBuilder struct {
	Where string
	Args  []interface{}
}

func Where(sql string, args ...interface{}) WhereBuilder {
	return WhereBuilder{
		Where: sql,
		Args:  args,
	}
}

func (w WhereBuilder) Build(db *gorm.DB) *gorm.DB {
	return db.Where(w.Where, w.Args...)
}

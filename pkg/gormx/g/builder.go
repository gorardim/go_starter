package g

import "gorm.io/gorm"

type Builder interface {
	Build(db *gorm.DB) *gorm.DB
}

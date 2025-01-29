package g

import "gorm.io/gorm"

var _ Builder = (*OrderBuilder)(nil)

type OrderBuilder struct {
	Order string
}

func Order(order string) OrderBuilder {
	return OrderBuilder{
		Order: order,
	}
}

func (o OrderBuilder) Build(db *gorm.DB) *gorm.DB {
	return db.Order(o.Order)
}

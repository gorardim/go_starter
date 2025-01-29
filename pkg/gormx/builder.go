package gormx

import "gorm.io/gorm"

type operator func(db *gorm.DB) *gorm.DB

type Builder struct {
	operators []operator
}

func (b *Builder) Where(query string, args ...interface{}) *Builder {
	b.operators = append(b.operators, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return b
}

func (b *Builder) Select(fields string) *Builder {
	b.operators = append(b.operators, func(db *gorm.DB) *gorm.DB {
		return db.Select(fields)
	})
	return b
}

func (b *Builder) WhereIf(predicate bool, query string, args ...interface{}) *Builder {
	if predicate {
		b.Where(query, args...)
	}
	return b
}

func (b *Builder) Order(order string) *Builder {
	b.operators = append(b.operators, func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	})
	return b
}

func (b *Builder) Group(group string) *Builder {
	b.operators = append(b.operators, func(db *gorm.DB) *gorm.DB {
		return db.Group(group)
	})
	return b
}

func (b *Builder) Having(query string, args ...interface{}) *Builder {
	b.operators = append(b.operators, func(db *gorm.DB) *gorm.DB {
		return db.Having(query, args...)
	})
	return b
}

func (b *Builder) Limit(limit int) *Builder {
	b.operators = append(b.operators, func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	})
	return b
}

func (b *Builder) Offset(offset int) *Builder {
	b.operators = append(b.operators, func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	})
	return b
}

func (b *Builder) Page(page, pageSize int) *Builder {
	if page < 1 {
		page = 1
	}
	b.operators = append(b.operators, func(db *gorm.DB) *gorm.DB {
		return db.Limit(pageSize).Offset((page - 1) * pageSize)
	})
	return b
}

func Where(query string, args ...interface{}) *Builder {
	b := &Builder{}
	b.Where(query, args...)
	return b
}

func Select(fields string) *Builder {
	b := &Builder{}
	b.Select(fields)
	return b
}

func WhereIf(predicate bool, query string, args ...interface{}) *Builder {
	if predicate {
		return Where(query, args...)
	} else {
		return &Builder{}
	}
}

func Order(order string) *Builder {
	b := &Builder{}
	b.Order(order)
	return b
}

func (b *Builder) Build(db *gorm.DB) *gorm.DB {
	for _, op := range b.operators {
		db = op(db)
	}
	return db
}

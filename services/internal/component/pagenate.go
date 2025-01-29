package component

import (
	"context"

	"app/pkg/sqx"
	"gorm.io/gorm"
)

type SimplePage[T any] struct {
	Items   []*T
	HasMore string
}

func SimplePaginate[T any](ctx context.Context, db *gorm.DB, page int, builder *sqx.Builder) (*SimplePage[T], error) {
	if page <= 0 {
		page = 1
	}
	clone := builder.Clone()
	list := make([]*T, 0)
	pageSize := 10
	query, args := clone.Limit(pageSize).Offset((page - 1) * pageSize).Build()
	if err := db.WithContext(ctx).Raw(query, args...).Scan(&list).Error; err != nil {
		return nil, err
	}

	hasMore := "N"
	if len(list) >= pageSize {
		hasMore = "Y"
	}
	return &SimplePage[T]{
		Items:   list,
		HasMore: hasMore,
	}, nil
}

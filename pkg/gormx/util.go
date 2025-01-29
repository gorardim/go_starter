package gormx

import (
	"context"
	"fmt"

	"app/pkg/sqx"
	"app/pkg/utils"

	"golang.org/x/exp/constraints"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Page[T any] struct {
	Total int64 `json:"total"`
	List  []*T  `json:"list"`
	Page  int   `json:"page"`
}

func Paginate[T any](ctx context.Context, db *gorm.DB, builder *sqx.Builder, page, pageSize int) (*Page[T], error) {
	clone := builder.Clone()
	if page < 1 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	result := &Page[T]{}
	list := make([]*T, 0)
	query, args := builder.Limit(pageSize).Offset((page - 1) * pageSize).Build()
	if err := db.WithContext(ctx).Raw(query, args...).Scan(&list).Error; err != nil {
		return nil, err
	}
	result.List = list
	// count
	var count int64
	query, args = clone.Select("count(*)").OrderBy("").Build()
	if err := db.WithContext(ctx).Raw(query, args...).Scan(&count).Error; err != nil {
		return nil, err
	}
	result.Total = count
	return result, nil
}

func FindAll[T any](ctx context.Context, db *gorm.DB, builder *sqx.Builder) ([]*T, error) {
	sql, args := builder.Build()
	list := make([]*T, 0)
	if err := db.WithContext(ctx).Raw(sql, args...).Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func Find[T any](ctx context.Context, db *gorm.DB, builder *sqx.Builder) (*T, error) {
	sql, args := builder.Limit(1).Build()
	var item T
	if err := db.WithContext(ctx).Raw(sql, args...).Scan(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func JsonPatch(v JsonObjectData) clause.Expr {
	return gorm.Expr(fmt.Sprintf(`JSON_MERGE_PATCH(%s, '%s')`, v.FieldName(), utils.JsonEncode(v)))
}

func Exists(ctx context.Context, db *gorm.DB, builder *sqx.Builder) (bool, error) {
	clone := builder.Clone()
	sql, args := clone.Select("count(*) as n").Limit(1).Build()
	var count int64
	if err := db.WithContext(ctx).Raw(sql, args...).Scan(&count).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return count > 0, nil
}

func FindAllBySql[T any](ctx context.Context, db *gorm.DB, sql string, args ...any) ([]*T, error) {
	list := make([]*T, 0)
	if err := db.WithContext(ctx).Raw(sql, args...).Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func Sum[T constraints.Float | constraints.Integer](ctx context.Context, db *gorm.DB, builder *sqx.Builder) (T, error) {
	sql, args := builder.Limit(1).Build()
	var sum T
	if err := db.WithContext(ctx).Raw(sql, args...).Scan(&sum).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return sum, nil
		}

		return sum, err
	}

	return sum, nil
}

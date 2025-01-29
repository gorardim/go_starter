package gormx

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"app/pkg/gormx/g"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type model interface {
	TableName() string
	PK() string
}

type Repo[T any] struct {
	db        *gorm.DB
	pk        string
	tableName string
}

func NewRepo[T any](db *gorm.DB) *Repo[T] {
	var t T
	m := any(&t).(model)
	return &Repo[T]{
		db:        db,
		pk:        m.PK(),
		tableName: m.TableName(),
	}
}

func (r *Repo[T]) Exists(ctx context.Context, query string, args ...any) (exists bool, err error) {
	var t T
	if err = r.db.WithContext(ctx).Where(query, args...).Take(&t).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *Repo[T]) Get(ctx context.Context, builder ...g.Builder) (*T, error) {
	var t T
	db := r.db.WithContext(ctx)
	for _, b := range builder {
		db = b.Build(db)
	}
	if err := db.Take(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Repo[T]) GetById(ctx context.Context, id int) (*T, error) {
	var t T
	if err := r.db.WithContext(ctx).Where(r.pk+" = ?", id).Take(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Repo[T]) Find(ctx context.Context, query string, args ...any) (*T, error) {
	var t T
	db := r.db.WithContext(ctx)
	if err := db.Where(query, args...).Take(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Repo[T]) FindBy(ctx context.Context, builder *Builder) (*T, error) {
	var t T
	db := builder.Build(r.db.WithContext(ctx))
	if err := db.Take(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Repo[T]) FindById(ctx context.Context, id int) (*T, error) {
	var t T
	if err := r.db.WithContext(ctx).Where(r.pk+" = ?", id).Take(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Repo[T]) FindAll(ctx context.Context, query string, args ...any) ([]*T, error) {
	var t []*T
	if err := r.db.WithContext(ctx).Where(query, args...).Find(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

func (r *Repo[T]) FindAllWithoutRef(ctx context.Context, query string, args ...any) ([]T, error) {
	var t []T
	if err := r.db.WithContext(ctx).Where(query, args...).Find(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

func (r *Repo[T]) FindAllBy(ctx context.Context, builder *Builder) ([]*T, error) {
	var t []*T
	db := builder.Build(r.db.WithContext(ctx))
	if err := db.Find(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

func (r *Repo[T]) Create(ctx context.Context, m *T) error {
	return r.TxCreate(ctx, r.db, m)
}

func (r *Repo[T]) TxCreate(ctx context.Context, tx *gorm.DB, m *T) error {
	return tx.WithContext(ctx).Create(m).Error
}

func (r *Repo[T]) CreateBatchTx(ctx context.Context, tx *gorm.DB, m []*T) error {
	if len(m) == 0 {
		return nil
	}
	return tx.WithContext(ctx).CreateInBatches(m, len(m)).Error
}

func (r *Repo[T]) Update(ctx context.Context, m *T, where string, args ...any) (int64, error) {
	return r.TxUpdate(ctx, r.db, m, where, args...)
}

func (r *Repo[T]) TxUpdate(ctx context.Context, tx *gorm.DB, m *T, where string, args ...any) (int64, error) {
	result := tx.WithContext(ctx).Where(where, args...).Updates(m)
	return result.RowsAffected, result.Error
}

func (r *Repo[T]) UpdateById(ctx context.Context, m *T) (int64, error) {
	return r.TxUpdateById(ctx, r.db, m)
}

func (r *Repo[T]) TxUpdateById(ctx context.Context, tx *gorm.DB, m *T) (int64, error) {
	result := tx.WithContext(ctx).Updates(m)
	return result.RowsAffected, result.Error
}

func (r *Repo[T]) Delete(ctx context.Context, where string, args ...any) (int64, error) {
	var model T
	tx := r.db.WithContext(ctx).Where(where, args...).Delete(&model)
	return tx.RowsAffected, tx.Error
}

func (r *Repo[T]) TxDelete(ctx context.Context, tx *gorm.DB, where string, args ...any) (int64, error) {
	var model T
	db := tx.WithContext(ctx).Where(where, args...).Delete(&model)
	return db.RowsAffected, db.Error
}

func (r *Repo[T]) UpdateMap(ctx context.Context, m map[string]any, where string, args ...any) (int64, error) {
	return r.TxUpdateMap(ctx, r.db, m, where, args...)
}

func (r *Repo[T]) TxUpdateMap(ctx context.Context, tx *gorm.DB, m map[string]any, where string, args ...any) (int64, error) {
	var model T
	result := tx.WithContext(ctx).Model(&model).Where(where, args...).Updates(m)
	return result.RowsAffected, result.Error
}

func (r *Repo[T]) Paginate(ctx context.Context, builder *Builder, page, pageSize int) (*Page[T], error) {
	if page < 1 {
		page = 1
	}
	result := &Page[T]{}
	result.Page = page
	// count
	count, err := r.CountBy(ctx, builder)
	if err != nil {
		return nil, err
	}
	list, err := r.FindAllBy(ctx, builder.Limit(pageSize).Offset((page-1)*pageSize))
	if err != nil {
		return nil, err
	}
	result.List = list

	result.Total = count
	return result, nil
}

func (r *Repo[T]) Count(ctx context.Context, where string, args ...any) (int64, error) {
	var count int64
	var model T
	if err := r.db.WithContext(ctx).Model(&model).Where(where, args...).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repo[T]) CountBy(ctx context.Context, builder *Builder) (int64, error) {
	var count int64
	var model T
	if err := builder.Build(r.db.WithContext(ctx)).Model(&model).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repo[T]) DB() *gorm.DB {
	return r.db
}

func (r *Repo[T]) Sum(ctx context.Context, filed string, where string, args ...any) (float64, error) {
	var sum float64
	if err := r.db.WithContext(ctx).Raw(fmt.Sprintf("SELECT IFNULL(SUM(%s),0) FROM `%s` WHERE %s", filed, r.tableName, where), args...).Scan(&sum).Error; err != nil {
		return 0, err
	}
	return sum, nil
}

func (r *Repo[T]) Avg(ctx context.Context, filed string, where string, args ...any) (float64, error) {
	var avg float64
	if err := r.db.WithContext(ctx).Raw(fmt.Sprintf("SELECT IFNULL(AVG(%s),0) FROM `%s` WHERE %s", filed, r.tableName, where), args...).Scan(&avg).Error; err != nil {
		return 0, err
	}
	return avg, nil
}

func (r *Repo[T]) SumN(ctx context.Context, fields []string, where string, args ...any) ([]float64, error) {
	buf := &bytes.Buffer{}
	for _, v := range fields {
		buf.WriteString(fmt.Sprintf("IFNULL(SUM(%s),0)", v))
		buf.WriteString(",")
	}
	buf.Truncate(buf.Len() - 1)

	rows, err := r.db.WithContext(ctx).Raw(fmt.Sprintf("SELECT %s FROM `%s` WHERE %s", buf.String(), r.tableName, where), args...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]interface{}, len(fields))
	for i, _ := range fields {
		result[i] = new(float64)
	}
	if rows.Next() {
		if err = rows.Scan(result...); err != nil {
			return nil, err
		}
	}
	var res []float64
	for _, v := range result {
		res = append(res, *(v.(*float64)))
	}
	return res, nil
}

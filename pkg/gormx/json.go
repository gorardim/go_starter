package gormx

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"app/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JsonObjectData interface {
	FieldName() string
}

func NewJsonObject[T JsonObjectData](data T) JsonObject[T] {
	return JsonObject[T]{
		Data: data,
	}
}

type JsonObject[T JsonObjectData] struct {
	Data T
}

func (j *JsonObject[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Data)
}

func (j *JsonObject[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &j.Data)
}

func (j *JsonObject[T]) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch value.(type) {
	case string, clause.Expr:
		return nil
	default:
		return json.Unmarshal(value.([]byte), &j.Data)
	}
}

func (j JsonObject[T]) Value() (driver.Value, error) {
	return json.Marshal(j.Data)
}

func (j JsonObject[T]) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	fieldName := j.Data.FieldName()
	if fieldName == "" {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{utils.JsonEncode(j.Data)},
		}
	}
	return clause.Expr{
		SQL:  fmt.Sprintf(`JSON_MERGE_PATCH(%s, ?)`, fieldName),
		Vars: []interface{}{utils.JsonEncode(j.Data)},
	}
}

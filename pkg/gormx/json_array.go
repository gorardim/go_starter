package gormx

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm/clause"
)

type JsonArray[T any] struct {
	Data []T
}

func NewJsonArray[T any](data []T) JsonArray[T] {
	return JsonArray[T]{
		Data: data,
	}
}

func (j *JsonArray[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Data)
}

func (j *JsonArray[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &j.Data)
}

func (j *JsonArray[T]) Scan(value interface{}) error {
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

func (j JsonArray[T]) Value() (driver.Value, error) {
	return json.Marshal(j.Data)
}

func (j JsonArray[T]) GormValue() clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{j.Data},
	}
}

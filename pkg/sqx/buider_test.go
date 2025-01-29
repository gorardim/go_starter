package sqx

import (
	"testing"
)

func TestSelect(t *testing.T) {
	query, args := Select("*").
		From("user u left join user_info ui on u.id = ui.user_id").
		Where("u.id = ?", 1).
		Where("u.name = ?", "hq").
		OrderBy("u.id desc").
		GroupBy("u.id").
		Having("u.id = ?", 1).
		Limit(10).
		Build()
	t.Log(query)
	t.Log(args)
}

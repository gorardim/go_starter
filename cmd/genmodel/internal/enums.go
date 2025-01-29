package internal

import (
	"fmt"
	"strings"
)

type Enum struct {
	Items      []EnumItem
	NamePrefix string
}

type EnumItem struct {
	Name    string
	Value   string
	Comment string
}

func parserEnums(options *Options, table string, columns []Column) []Enum {
	enums := make([]Enum, 0)
	for _, column := range columns {
		if column.Comment == "" {
			continue
		}
		name := fmt.Sprintf("%s.%s", table, column.Field)

		if !options.enumTypesMap[name] && !checkFieldName(column.Field) {
			continue
		}
		items := parseEnum(column.Comment)
		if len(items) == 0 {
			continue
		}
		enums = append(enums, Enum{
			Items:      items,
			NamePrefix: fmt.Sprintf("%s%s", clsName(table), clsName(column.Field)),
		})
	}
	return enums
}

func parseEnum(first string) []EnumItem {
	// 业务类型:投资INVEST,推荐RECOMMEND,团队TEAM
	idx := strings.Index(first, ":")
	sp := strings.Split(strings.TrimSpace(first[idx+1:]), ",")

	r := make([]EnumItem, 0)
	for _, s := range sp {
		s1 := strings.TrimSpace(s)
		if s1 == "" {
			continue
		}
		w1, w2 := splitWord(s1)
		if w2 == "" {
			continue
		}
		if w1 == "" {
			w1 = w2
		}
		r = append(r, EnumItem{
			Name:    clsName(strings.ToLower(w2)),
			Value:   w2,
			Comment: w1,
		})
	}
	return r
}

func splitWord(s1 string) (string, string) {
	idx := len(s1)
	// 分割中英文
	for i, r := range s1 {
		// is A-Z a-z
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			idx = i
			break
		}
	}
	return s1[:idx], s1[idx:]
}

func checkFieldName(name string) bool {
	name = strings.ToLower(name)
	for _, s := range []string{"type", "status"} {
		if strings.Contains(name, s) {
			return true
		}
	}
	return false
}

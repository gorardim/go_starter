package sqx

import (
	"bytes"
	"strconv"
)

const (
	OpeSelect = "SELECT"
	OpeInsert = "INSERT"
	OpeUpdate = "UPDATE"
	OpeDelete = "DELETE"
)

type Filter struct {
	Query string
	Args  []any
}

func NewFilter(query string, args ...any) Filter {
	return Filter{
		Query: query,
		Args:  args,
	}
}

type Builder struct {
	// 操作
	operate string
	fields  string
	from    []Filter
	where   []Filter
	order   string
	group   string
	having  Filter
	limit   int
	offset  int
}

// Select 生成查询语句
func (s *Builder) Select(fields string) *Builder {
	s.operate = OpeSelect
	s.fields = fields
	return s
}

// From 表
func (s *Builder) From(from string, args ...any) *Builder {
	s.from = append(s.from, Filter{
		Query: from,
		Args:  args,
	})
	return s
}

func (s *Builder) LeftJoin(from string, args ...any) *Builder {
	return s.From("LEFT JOIN "+from, args...)
}

func (s *Builder) InnerJoin(from string, args ...any) *Builder {
	return s.From("INNER JOIN "+from, args...)
}

// Where 条件
func (s *Builder) Where(query string, args ...any) *Builder {
	s.where = append(s.where, Filter{
		Query: query,
		Args:  args,
	})
	return s
}

// WhereIf 条件
func (s *Builder) WhereIf(predicate bool, query string, args ...any) *Builder {
	if predicate {
		s.Where(query, args...)
	}
	return s
}

func (s *Builder) WhereIfFunc(predicate bool, fn func() Filter) *Builder {
	if predicate {
		filter := fn()
		s.Where(filter.Query, filter.Args...)
	}
	return s
}

func (s *Builder) WhereNotEmpty(predicate string, query string, args ...any) *Builder {
	return s.WhereIf(predicate != "", query, args...)
}

// GroupBy 分组
func (s *Builder) GroupBy(group string) *Builder {
	s.group = group
	return s
}

// Having 分组条件
func (s *Builder) Having(query string, args ...any) *Builder {
	s.having = Filter{
		Query: query,
		Args:  args,
	}
	return s
}

// WhereHaving 分组条件
func (s *Builder) WhereHaving(predicate bool, query string, args ...any) *Builder {
	if predicate {
		s.having = Filter{
			Query: query,
			Args:  args,
		}
	}

	return s
}

// OrderBy 排序
func (s *Builder) OrderBy(order string) *Builder {
	s.order = order
	return s
}

func (s *Builder) Limit(limit int) *Builder {
	s.limit = limit
	return s
}

func (s *Builder) Offset(offset int) *Builder {
	s.offset = offset
	return s
}

func (s *Builder) Page(page int, pageSize int) *Builder {
	if page <= 0 {
		page = 1
	}
	s.limit = pageSize
	s.offset = (page - 1) * pageSize
	return s
}

// Clone copy builder
func (s *Builder) Clone() *Builder {
	return &Builder{
		operate: s.operate,
		fields:  s.fields,
		from:    s.from,
		where:   s.where,
		order:   s.order,
		group:   s.group,
		having:  s.having,
		limit:   s.limit,
		offset:  s.offset,
	}
}

func (s *Builder) Build() (string, []any) {
	buf := bytes.Buffer{}
	args := make([]any, 0)
	switch s.operate {
	case OpeSelect:
		buf.WriteString(s.operate)
		buf.WriteString(" ")
		buf.WriteString(s.fields)
		buf.WriteString(" FROM ")
		for _, item := range s.from {
			buf.WriteString(item.Query)
			buf.WriteString(" ")
			args = append(args, item.Args...)
		}

		if len(s.where) > 0 {
			buf.WriteString(" WHERE ")
			for i, v := range s.where {
				if i > 0 {
					buf.WriteString(" AND ")
				}
				buf.WriteString("(")
				buf.WriteString(v.Query)
				buf.WriteString(")")
				args = append(args, v.Args...)
			}
		}
		if s.group != "" {
			buf.WriteString(" GROUP BY ")
			buf.WriteString(s.group)
		}
		if s.having.Query != "" {
			buf.WriteString(" HAVING ")
			buf.WriteString(s.having.Query)
			args = append(args, s.having.Args...)
		}
		if s.order != "" {
			buf.WriteString(" ORDER BY ")
			buf.WriteString(s.order)
		}
		if s.limit > 0 {
			buf.WriteString(" LIMIT ")
			buf.WriteString(strconv.Itoa(s.limit))
		}
		if s.offset > 0 {
			buf.WriteString(" OFFSET ")
			buf.WriteString(strconv.Itoa(s.offset))
		}
	}
	return buf.String(), args
}

func (s *Builder) String() string {
	query, _ := s.Build()
	return query
}

func Select(fields string) *Builder {
	return &Builder{
		operate: OpeSelect,
		fields:  fields,
	}
}

func From(from string, args ...any) *Builder {
	return Select("*").From(from, args...)
}

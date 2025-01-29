package internal

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	"app/cmd/internal"
	"gorm.io/gorm"
)

//go:embed tpl.gohtml
var modelTemplate string
var modelTpl = template.Must(template.New("model").Funcs(map[string]any{
	"clsName": clsName,
	"typName": typName,
	"minus": func(a, b int) int {
		return a - b
	},
}).Parse(modelTemplate))

type Options struct {
	TypeMapping  map[string]string
	EnumTypes    []string
	enumTypesMap map[string]bool
}

func Generate(ops *Options) {
	ops.enumTypesMap = make(map[string]bool)
	for _, v := range ops.EnumTypes {
		ops.enumTypesMap[v] = true
	}
	db, err := internal.OpenDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	for _, table := range getAllTable(db) {
		if strings.HasPrefix(table, "e_") {
			continue
		}
		genModel(ops, db, table)
	}

}

type Data struct {
	Table   string
	Comment string
	Columns []Column
	PK      string
	Enums   []Enum
}

func genModel(options *Options, db *gorm.DB, table string) {
	filename := fmt.Sprintf("./api/model/%s.go", table)
	if fileExists(filename) {
		return
	}
	columns := getTableColumns(options, db, table)
	log.Printf("write file: %s", filename)

	buf := &bytes.Buffer{}
	err := modelTpl.Execute(buf, Data{
		Table:   table,
		Comment: getTableStatus(db, table).Comment,
		Columns: columns,
		PK:      getPK(columns),
		Enums:   parserEnums(options, table, columns),
	})
	if err != nil {
		log.Fatalf("failed to execute template: %s", err)
	}
	// format code
	source, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("failed to format source: %s", err)
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	_, err = file.Write(source)
	if err != nil {
		log.Fatalf("failed to write source: %s", err)
	}
}

func getAllTable(db *gorm.DB) []string {
	tables := make([]string, 0)
	db.Raw("show tables").Scan(&tables)
	return tables
}

type Column struct {
	Field      string `gorm:"column:Field"`
	Type       string `gorm:"column:Type"`
	Null       string `gorm:"column:Null"`
	Key        string `gorm:"column:Key"`
	Default    string `gorm:"column:Default"`
	Extra      string `gorm:"column:Extra"`
	Comment    string `gorm:"column:Comment"`
	Table      string `gorm:"-"`
	CustomType string `gorm:"-"`
}

func getTableColumns(options *Options, db *gorm.DB, table string) []Column {
	columns := make([]Column, 0)
	db.Raw("show full columns from " + table).Scan(&columns)
	for i := range columns {
		columns[i].Table = table
		name := fmt.Sprintf("%s.%s", table, columns[i].Field)
		if typ, ok := options.TypeMapping[name]; ok {
			columns[i].CustomType = typ
		}
	}
	return columns
}

type TableStatus struct {
	Name    string
	Comment string
}

func getTableStatus(db *gorm.DB, table string) TableStatus {
	var s TableStatus
	db.Raw("show table status where name = ?", table).Scan(&s)
	return s
}

func clsName(s string) string {
	return came(s, true)
}

func came(s string, firstToUpper bool) string {
	up := firstToUpper
	var r string
	for _, v := range s {
		if v == '_' {
			up = true
			continue
		}

		if up {
			// check v is letter
			if v >= 'a' && v <= 'z' {
				v = v - 'a' + 'A'
			}
			up = false
		}
		r += string(v)
	}
	return r
}

func typName(c Column) string {
	if c.CustomType != "" {
		return c.CustomType
	}
	s := c.Type
	switch {
	case strings.Contains(s, "int"):
		return "int"
	case strings.Contains(s, "float"), strings.Contains(s, "double"), strings.Contains(s, "decimal"):
		return "float64"
	case strings.Contains(s, "date"), strings.Contains(s, "time"):
		return "time.Time"
	default:
		return "string"
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func getPK(columns []Column) string {
	for _, column := range columns {
		if column.Key == "PRI" {
			return column.Field
		}
	}
	return ""
}

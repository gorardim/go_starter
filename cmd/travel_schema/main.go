package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"app/cmd/internal"
	"gorm.io/gorm"
)

var table string

type TableSchema struct {
	Table       string `gorm:"column:Table"`
	CreateTable string `gorm:"column:Create Table"`
}

// 写入schema.sql文件
func main() {
	flag.StringVar(&table, "table", "", "table name")
	flag.Parse()
	if table == "" {
		log.Fatal("table is empty")
	}
	tableList := strings.Split(table, ",")
	db, err := internal.OpenDB()
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %v", err))
	}
	buf := bytes.Buffer{}
	// 获取创建表的sql语句
	for _, t := range tableList {
		tableName := strings.TrimSpace(t)
		if tableName == "" {
			continue
		}
		split := strings.Split(tableName, ":")
		tableName = split[0]
		tag := ""
		if len(split) > 1 {
			tag = split[1]
		}
		var tableSchema TableSchema
		if err = db.Raw("show create table " + tableName).Scan(&tableSchema).Error; err != nil {
			log.Fatalf("show create table %s failed: %v", tableName, err)
		}
		buf.WriteString(tableSchema.CreateTable)
		buf.WriteString(";\n\n")

		// 数据集,写入最新的n条数据，如果不存在则mock一条
		if tag != "none" {
			genFixture(db, tableName, tag)
		}
	}
	fmt.Printf("write schema.sql\n")
	// 替换schema.sql文件
	if err = os.WriteFile("schema.sql", buf.Bytes(), 0644); err != nil {
		log.Fatalf("write schema.sql failed: %v", err)
	}
}

type Column struct {
	Field string `gorm:"column:Field"`
	Type  string `gorm:"column:Type"`
	Key   string `gorm:"column:Key"`
}

func genFixture(db *gorm.DB, tableName string, tag string) {
	// 获取表的字段结构
	var columns []Column
	if err := db.Raw(fmt.Sprintf("show full columns from %s", tableName)).Scan(&columns).Error; err != nil {
		log.Fatalf("show full columns from %s failed: %v", tableName, err)
	}

	// pk
	var pk string
	for _, c := range columns {
		if strings.Contains(c.Key, "PRI") {
			pk = c.Field
			break
		}
	}

	if pk == "" {
		log.Fatalf("table %s has no primary key", tableName)
	}

	// 获取表的数据
	var data []map[string]interface{}

	limit := 10
	switch tag {
	case "export_all":
		limit = 100000
	default:
		// is number
		if n, err := strconv.Atoi(tag); err == nil {
			limit = n
		}
	}

	if err := db.Raw(fmt.Sprintf("select * from %s order by `%s` desc limit %d", tableName, pk, limit)).Scan(&data).Error; err != nil {
		log.Fatalf("select * from %s failed: %v", tableName, err)
	}

	// check fixtures exists
	if _, err := os.Stat("fixtures"); os.IsNotExist(err) {
		if err = os.Mkdir("fixtures", 0755); err != nil {
			log.Fatalf("mkdir fixtures failed: %v", err)
		}
	}

	// 生成fixture
	if len(data) > 0 {
		out := marshalYmlOrdered(data, columns)
		fmt.Printf("write fixtures/%s.yml\n", tableName)
		if err := os.WriteFile(fmt.Sprintf("fixtures/%s.yml", tableName), out, 0644); err != nil {
			log.Fatalf("write fixtures/%s.yml failed: %v", tableName, err)
		}
	} else {
		// mock一条数据
		mockData := make(map[string]interface{})
		for _, c := range columns {
			mockData[c.Field] = mockValue(c)
		}
		out := marshalYmlOrdered([]map[string]interface{}{mockData}, columns)
		fmt.Printf("write fixtures/%s.yml\n", tableName)
		if err := os.WriteFile(fmt.Sprintf("fixtures/%s.yml", tableName), out, 0644); err != nil {
			log.Fatalf("write fixtures/%s.yml failed: %v", tableName, err)
		}
	}
}

func mockValue(col Column) any {
	switch {
	case strings.HasPrefix(col.Type, "varchar"):
		return col.Field
	case strings.HasPrefix(col.Type, "int"):
		return 1
	case strings.HasPrefix(col.Type, "datetime"):
		return "2021-01-01 00:00:00"
	case strings.HasPrefix(col.Type, "timestamp"):
		return "2021-01-01 00:00:00"
	case strings.HasPrefix(col.Type, "decimal"):
		return 1.01
		// json
	case strings.HasPrefix(col.Type, "json"):
		return "{}"

	default:
		return col.Field
	}
}

func marshalYmlOrdered(data []map[string]interface{}, fields []Column) []byte {
	var buf bytes.Buffer
	for _, d := range data {
		for i, f := range fields {
			if i == 0 {
				buf.WriteString("- ")
			} else {
				buf.WriteString("  ")
			}
			value := d[f.Field]
			switch v := value.(type) {
			case string:
				value = strconv.Quote(v)
				if f.Type == "json" {
					value = fmt.Sprintf("|\n    %s", v)
				}
			case time.Time:
				value = v.Format("2006-01-02 15:04:05")
			}
			fmtValue := value
			if value == nil {
				fmtValue = "null"
			}
			buf.WriteString(fmt.Sprintf("%s: %v\n", f.Field, fmtValue))
		}
		buf.WriteString("\n")
	}
	return buf.Bytes()
}

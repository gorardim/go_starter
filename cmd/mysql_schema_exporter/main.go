package main

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strings"

	"app/cmd/internal"
	"gorm.io/gorm"
)

var autoIncrementReg = regexp.MustCompile(`AUTO_INCREMENT=\d+`)

func main() {
	db, err := internal.OpenDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	tables := getAllTable(db)
	buf := &bytes.Buffer{}
	for _, table := range tables {
		// write drop table
		buf.WriteString("DROP TABLE IF EXISTS " + table + ";\n")
		info := getCreateTable(db, table)
		info = strings.Replace(info, "ENGINE=InnoDB", "", -1)
		info = strings.Replace(info, "DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ", "", -1)
		info = strings.Replace(info, "DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ", "", -1)
		info = strings.Replace(info, "DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ", "", -1)
		// rm AUTO_INCREMENT
		info = string(autoIncrementReg.ReplaceAll([]byte(info), []byte("")))
		buf.WriteString(info)
		buf.WriteString(";\n\n")
	}
	err = os.WriteFile("./docs/schema.sql", buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
}

func getAllTable(db *gorm.DB) []string {
	tables := make([]string, 0)
	db.Raw("show tables").Scan(&tables)
	return tables
}

type createTable struct {
	Table       string `gorm:"column:Table"`
	CreateTable string `gorm:"column:Create Table"`
}

func getCreateTable(db *gorm.DB, table string) string {
	v := &createTable{}
	if err := db.Raw("show create table " + table).Scan(v).Error; err != nil {
		log.Fatalf("failed to get create table: %v", err)
	}
	return v.CreateTable
}

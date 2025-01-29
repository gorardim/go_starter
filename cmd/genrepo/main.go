package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"text/template"

	"app/cmd/internal"

	"gorm.io/gorm"
)

//go:embed tpl.gohtml
var repoTemplate string
var repoTpl = template.Must(template.New("repo").Funcs(map[string]any{
	"clsName": clsName,
}).Parse(repoTemplate))

func main() {
	db, err := internal.OpenDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	tables := getAllTable(db)
	for _, table := range tables {
		genRepo(db, table)
	}
}

func genRepo(db *gorm.DB, table string) {
	filename := fmt.Sprintf("./services/internal/repo/%s.go", table)
	if fileExists(filename) {
		return
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fmt.Printf("write file: %s\n", filename)
	err = repoTpl.Execute(file, map[string]any{
		"Table":   table,
		"Comment": getTableStatus(db, table).Comment,
	})
	if err != nil {
		log.Fatalf("failed to execute template: %s", err)
	}
}

func getAllTable(db *gorm.DB) []string {
	tables := make([]string, 0)
	db.Raw("show tables").Scan(&tables)
	return tables
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
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

type TableStatus struct {
	Name    string
	Comment string
}

func getTableStatus(db *gorm.DB, table string) TableStatus {
	var s TableStatus
	db.Raw("show table status where name = ?", table).Scan(&s)
	return s
}

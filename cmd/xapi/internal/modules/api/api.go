package api

import (
	"log"
	"os"
	"path/filepath"

	"app/cmd/xapi/internal/parser"
)

func Run(scanDir string, title string) {
	doc, err := parser.ParseDir(scanDir)
	if err != nil {
		log.Fatalf("parse dir error: %v", err)
	}
	// save openapi
	docFilename := filepath.Join(scanDir, "openapi.yml")
	src := docToOpenApi(doc, title).Bytes()
	// merge
	appendFile := filepath.Join(scanDir, "append.openapi.yml")
	if fileExists(appendFile) {
		readFile, err := os.ReadFile(appendFile)
		if err != nil {
			log.Fatalf("read file error: %v", err)
		}
		src, err = MergeOpenApiYamlContent(src, readFile)
		if err != nil {
			log.Fatalf("merge file error: %v", err)
		}
	}
	err = os.WriteFile(docFilename, src, 0644)
	if err != nil {
		log.Fatalf("write file error: %v", err)
	}
	// generate http
	generateHttp(doc)
}

// check if file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

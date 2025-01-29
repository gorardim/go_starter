package parser

import (
	"strings"

	"app/cmd/xapi/internal/utils"
	"golang.org/x/tools/go/packages"
)

func ParseDir(dir string) (*Document, error) {
	doc := newDocument()
	load, err := packages.Load(&packages.Config{
		Mode: packages.LoadAllSyntax,
		Dir:  dir,
	})
	if err != nil {
		return nil, err
	}

	for _, pkg := range load {
		info := Info(pkg)
		for i, astFle := range pkg.Syntax {
			// 如果文件名称是 .xx.go 则跳过
			filename := pkg.GoFiles[i]
			if strings.Contains(strings.TrimSuffix(filename, ".go"), ".") {
				continue
			}
			ctx := &Context{
				document: doc,
				file:     astFle,
				pkg:      pkg,
				info:     info,
				stack:    &utils.NodeStack{},
				filename: filename,
			}
			if err = parse(ctx); err != nil {
				return nil, err
			}
		}
	}
	return doc, nil
}

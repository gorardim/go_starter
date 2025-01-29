package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
)

func parseExternalSchema(pkgDir string) []*Schema {
	dir, err := os.ReadDir(pkgDir)
	if err != nil {
		panic(err)
	}
	// read files
	var result []*Schema
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		filename := filepath.Join(pkgDir, file.Name())
		content, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		// parse
		astFile, err := parser.ParseFile(token.NewFileSet(), filename, content, parser.ParseComments)
		if err != nil {
			panic(err)
		}

		// read all struct
		for _, decl := range astFile.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				if genDecl.Tok == token.TYPE {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if structType, ok := typeSpec.Type.(*ast.StructType); ok {
								// parse struct
								schema := &Schema{}
								schema.Name = typeSpec.Name.Name
								for _, field := range structType.Fields.List {
									schemaField := &Field{}
									schemaField.Name = field.Names[0].Name
									typ := getTypeName(field)
									if typ == "" {
										goto NEXT
									}
									// 原生类型
									if !isNativeTypeOrImportType(typ) {
										typ = fmt.Sprintf("%s.%s", astFile.Name.Name, typ)
									}
									schemaField.Type = typ

									if field.Doc != nil {
										schemaField.Comment = lo.Map(field.Doc.List, func(item *ast.Comment, _ int) string {
											return item.Text
										})
									}
									schema.Fields = append(schema.Fields, schemaField)
								}
								result = append(result, schema)
							}
						}
					}
				}
			}
		NEXT:
		}
	}
	return result
}

func getTypeName(node ast.Node) string {
	switch v := node.(type) {
	case *ast.Field:
		return getTypeName(v.Type)
	case *ast.Ident:
		return v.Name
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", getTypeName(v.Elt))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", v.X.(*ast.Ident).Name, v.Sel.Name)
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", getTypeName(v.X))
	case *ast.IndexExpr:
		return fmt.Sprintf("%s[%s]", getTypeName(v.X), getTypeName(v.Index))
	default:
		fmt.Printf("unknown type: %T\n", v)
		return ""
	}
}

func isNativeTypeOrImportType(name string) bool {
	if strings.Contains(name, ".") {
		return true
	}
	switch name {
	case "string", "int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "float32", "float64", "bool", "byte", "rune":
		return true
	default:
		return false
	}
}

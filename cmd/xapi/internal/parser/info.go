package parser

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

type PackageInfo struct {
	StructMap map[string]*StructInfo
}

func (i PackageInfo) Struct(of types.Object) *StructInfo {
	return i.StructMap[fmt.Sprintf("%s.%s", of.Pkg().Path(), of.Name())]
}

type StructInfo struct {
	Fields   []*ast.Field
	Name     string
	Comments []*ast.Comment
}

func Info(pkg *packages.Package) *PackageInfo {
	info := PackageInfo{
		StructMap: make(map[string]*StructInfo),
	}

	for _, astFile := range pkg.Syntax {
		// read all struct
		ast.Inspect(astFile, func(node ast.Node) bool {
			if node == nil {
				return true
			}
			switch v := node.(type) {
			case *ast.StructType:
				path, _ := astutil.PathEnclosingInterval(astFile, v.Pos(), v.Pos())
				structInfo := StructInfo{}
				for _, pathNode := range path {
					switch pathNode.(type) {
					case *ast.TypeSpec:
						typeSpec := pathNode.(*ast.TypeSpec)
						structInfo.Name = typeSpec.Name.Name
					case *ast.GenDecl:
						genDecl := pathNode.(*ast.GenDecl)
						if genDecl.Doc != nil {
							structInfo.Comments = genDecl.Doc.List
						}
					case *ast.StructType:
						structInfo.Fields = v.Fields.List
					}
				}
				info.StructMap[fmt.Sprintf("%s.%s", pkg.ID, structInfo.Name)] = &structInfo
			}
			return true
		})
	}

	return &info
}

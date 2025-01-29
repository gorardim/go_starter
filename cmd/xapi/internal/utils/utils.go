package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func ParseFile(fileName string) (*ast.File, error) {
	// read file
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	// parse file
	fileSet := token.NewFileSet()
	return parser.ParseFile(fileSet, "", content, parser.ParseComments)
}

func ParseTag(value string) map[string]string {
	value = value[1 : len(value)-1]
	tags := make(map[string]string)
	isReadKey := true
	key := ""
	val := ""
	for _, c := range value {
		switch isReadKey {
		case true:
			if unicode.IsSpace(c) || c == ':' {
				continue
			}
			if c == '"' {
				isReadKey = false
				continue
			}
			key += string(c)
		case false:
			if c == '"' {
				isReadKey = true
				tags[key] = val
				val = ""
				key = ""
				continue
			}
			val += string(c)
		}
	}
	return tags
}

// 获取注释
func GetNodeComment(node ast.Node) []string {
	results := make([]string, 0)
	var doc *ast.CommentGroup
	switch v := node.(type) {
	case *ast.GenDecl:
		doc = v.Doc
	case *ast.Field:
		doc = v.Doc
	}
	if doc != nil {
		for _, comment := range doc.List {
			results = append(results, comment.Text)
		}
	}
	return results
}

func GetIdentName(node ast.Node) string {
	switch v := node.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.TypeSpec:
		return v.Name.Name
	}
	return ""
}

func GetType(node ast.Node) string {
	switch v := node.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.TypeSpec:
		return v.Name.Name
	case *ast.StarExpr:
		return GetType(v.X)
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", v.X.(*ast.Ident).Name, v.Sel.Name)
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", GetType(v.Elt))
	}
	return ""
}

func IsContextExpr(node ast.Expr) bool {
	expr, ok := node.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	return expr.Sel.Name == "Context"
}

func ReadDir(path string) ([]string, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var r []string
	for _, dir := range dirs {
		if dir.IsDir() || !strings.HasSuffix(dir.Name(), ".go") || strings.Count(dir.Name(), ".") != 1 {
			continue
		}
		r = append(r, filepath.Join(path, dir.Name()))
	}
	return r, nil
}

func CameCase(s string) string {
	var r string
	var up bool
	for i, c := range s {
		if i == 0 {
			r += strings.ToLower(string(c))
			continue
		}
		if c == '_' || c == '-' {
			up = true
			continue
		}
		if up {
			r += strings.ToUpper(string(c))
			up = false
		} else {
			r += string(c)
		}
	}
	return r
}

// ParseDirective 解析指令
// 指令格式: //x:<name> param1 param2 param3 ...
func ParseDirective(directive string) (name string, params []string) {
	directive = strings.TrimSpace(directive)
	if !strings.HasPrefix(directive, "//x:") {
		return "", nil
	}
	directive = directive[4:]
	index := strings.Index(directive, " ")
	if index == -1 {
		return directive, nil
	}
	return directive[:index], strings.Split(directive[index+1:], " ")
}

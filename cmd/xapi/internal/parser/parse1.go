package parser

import (
	"go/ast"
	"regexp"

	"app/cmd/xapi/internal/utils"
	"golang.org/x/tools/go/packages"
)

type Context struct {
	document *Document
	file     *ast.File
	pkg      *packages.Package
	info     *PackageInfo
	stack    *utils.NodeStack
	filename string
}

func parse(ctx *Context) error {
	ast.Inspect(ctx.file, func(node ast.Node) bool {
		if node == nil {
			ctx.stack.Pop()
			return true
		}
		ctx.stack.Push(node)

		switch v := node.(type) {
		case *ast.InterfaceType:
			// 如果是nsq则解析nsq
			if nsqParse(ctx, v) {
				return true
			}
			parseService(ctx, v)
		case *ast.StructType:
			parseSchema(ctx, v)
		default:
			// fmt.Println(node)
		}
		return true
	})
	return nil
}

func parseSchema(ctx *Context, v *ast.StructType) {
	schema := &Schema{}
	schema.Name = utils.GetIdentName(ctx.stack.Top(1))
	if schema.Name == "" {
		return
	}
	schema.Fields = parseFields(v.Fields.List, ctx.pkg, ctx.info, ctx.document)
	ctx.document.Schemas[schema.Name] = schema
}

func parseService(ctx *Context, v *ast.InterfaceType) {
	service := &Service{
		FileName: ctx.filename,
		Package:  ctx.pkg.Name,
	}
	service.Comment = utils.GetNodeComment(ctx.stack.Top(2))
	service.Name = utils.GetIdentName(ctx.stack.Top(1))

	for _, field := range v.Methods.List {
		rpc := &Rpc{}
		funcType := field.Type.(*ast.FuncType)
		paramList := funcType.Params.List
		// check param list length
		if len(paramList) != 2 || !utils.IsContextExpr(paramList[0].Type) {
			continue
		}
		rpc.Name = field.Names[0].Name
		rpc.Comment = utils.GetNodeComment(field)

		for i, v := range rpc.Comment {
			api := parseXApi(v)
			if api != nil {
				rpc.Path = api[1]
				rpc.Method = api[0]
				rpc.Comment[i] = api[2]
				break
			}
		}
		if rpc.Path == "" {
			continue
		}
		rpc.Request = paramList[1].Type.(*ast.StarExpr).X.(*ast.Ident).Name
		// check results length
		if len(funcType.Results.List) != 2 {
			continue
		}
		rpc.Response = funcType.Results.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
		service.RpcArr = append(service.RpcArr, rpc)
	}
	if len(service.RpcArr) > 0 {
		ctx.document.Services = append(ctx.document.Services, service)
	}
}

func parseFields(astFields []*ast.Field, pkg *packages.Package, info *PackageInfo, document *Document) []*Field {
	var fields []*Field
	for _, fieldNode := range astFields {
		field := &Field{}
		if fieldNode.Names == nil {
			// anonymous field
			of := pkg.TypesInfo.ObjectOf(fieldNode.Type.(*ast.Ident))
			if of == nil {
				continue
			}
			structInfo := info.Struct(of)
			if structInfo != nil {
				fields = append(fields, parseFields(structInfo.Fields, pkg, info, document)...)
			}
			continue
		}
		field.Name = fieldNode.Names[0].Name
		if field.Name == "" {
			continue
		}
		field.Type = utils.GetType(fieldNode.Type)
		field.Comment = utils.GetNodeComment(fieldNode)
		// tag
		if fieldNode.Tag != nil && len(fieldNode.Tag.Value) > 0 {
			tag := utils.ParseTag(fieldNode.Tag.Value)
			if tag["json"] == "-" {
				continue
			}
			field.Tags = tag
		}
		fields = append(fields, field)
	}
	return fields
}

var apiReg = regexp.MustCompile(`^//x:api\s+(\w+)\s+(\S+)\s+(.*)$`)

// parse x:api
func parseXApi(content string) []string {
	// parse // x:api method path comment
	subMatch := apiReg.FindStringSubmatch(content)
	if subMatch == nil {
		return nil
	}
	return subMatch[1:]
}

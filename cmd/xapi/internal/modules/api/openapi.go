package api

import (
	"bytes"
	"strconv"
	"strings"

	"app/cmd/xapi/internal/parser"
)

func docToOpenApi(doc *parser.Document, title string) *bytes.Buffer {
	buf := &bytes.Buffer{}
	buf.WriteString("openapi: 3.0.0\n")
	buf.WriteString("info:\n")
	buf.WriteString("  title: " + title + "\n")
	buf.WriteString("  version: 1.0.0\n")

	// write tags
	buf.WriteString("tags:\n")
	for _, service := range doc.Services {
		buf.WriteString("  - name: " + service.Name + "\n")
		buf.WriteString("    description: " + strings.Replace(getComment(service.Comment), service.Name, "", 1) + "\n")
	}

	// write paths
	buf.WriteString("paths:\n")
	for _, service := range doc.Services {
		for _, rpc := range service.RpcArr {
			buf.WriteString("  " + rpc.Path + ":\n")
			buf.WriteString("    " + rpc.Method + ":\n")
			// description
			buf.WriteString("      description: " + getComment(rpc.Comment) + "\n")
			// tags
			buf.WriteString("      tags: [\"" + service.Name + "\"]\n")
			// summary
			buf.WriteString("      summary: " + getComment(rpc.Comment) + "\n")
			// operationId
			buf.WriteString("      operationId: " + service.Name + "_" + rpc.Name + "\n")

			// parameters
			buf.WriteString("      parameters:\n")
			buf.WriteString("        - in: header\n")
			buf.WriteString("          name: Accept-Language\n")
			buf.WriteString("          description: Accept Language\n")
			buf.WriteString("          schema:\n")
			buf.WriteString("            type: string\n")

			// requestBody
			buf.WriteString("      requestBody:\n")
			buf.WriteString("        content:\n")
			buf.WriteString("          application/json:\n")
			buf.WriteString("            schema:\n")
			buf.WriteString("              $ref: \"#/components/schemas/" + rpc.Request + "\"\n")
			// responses
			buf.WriteString("      responses:\n")
			buf.WriteString("        200:\n")
			buf.WriteString("          description: OK\n")
			buf.WriteString("          content:\n")
			buf.WriteString("            application/json:\n")
			buf.WriteString("              schema:\n")
			buf.WriteString("                $ref: \"#/components/schemas/" + rpc.Response + "\"\n")
		}
	}

	// components
	buf.WriteString("components:\n")
	buf.WriteString("  schemas:\n")
	for name, schema := range doc.Schemas {
		buf.WriteString("    " + name + ":\n")
		buf.WriteString("      type: object\n")
		buf.WriteString("      properties:\n")
		for _, field := range schema.Fields {
			buf.WriteString("        " + getOpenapiName(field) + ":\n")
			apiType := getOpenApiType(field.Type)

			switch apiType {
			case "array":
				buf.WriteString("          type: " + apiType + "\n")
				buf.WriteString("          items:\n")
				//  ref
				if isOpenApiBasicType(field.Type[2:]) {
					buf.WriteString("            type: " + getOpenApiType(field.Type[2:]) + "\n")
				} else {
					buf.WriteString("            $ref: \"#/components/schemas/" + field.Type[2:] + "\"\n")
				}
			case "object":
				buf.WriteString("          $ref: \"#/components/schemas/" + field.Type + "\"\n")
			default:
				buf.WriteString("          type: " + apiType + "\n")
				buf.WriteString("          description: " + getComment(field.Comment) + "\n")
			}
		}
	}
	return buf
}

func getOpenapiName(field *parser.Field) string {
	if field.Tags == nil || field.Tags["json"] == "" {
		return field.Name
	}
	return strings.Split(field.Tags["json"], ",")[0]
}

func getOpenApiType(typ string) string {
	switch typ {
	case "int", "int32", "uint", "uint32", "uint64":
		return "integer"
	case "int64", "float32", "float64":
		return "string"
	case "time.Time", "string":
		return "string"
	default:
		if strings.HasPrefix(typ, "[]") {
			return "array"
		}
		return "object"
	}
}

// 是否为openapi 基本类型
func isOpenApiBasicType(typ string) bool {
	switch typ {
	case "integer", "string":
		return true
	default:
		return false
	}
}

func getComment(comments []string) string {
	var r []string

	if len(comments) == 0 {
		return "\"\""
	}

	for _, comment := range comments {
		// replace \n to space
		comment = strings.Replace(comment, "\n", " ", -1)
		// trim //
		comment = strings.TrimPrefix(comment, "//")
		// trim space
		comment = strings.TrimSpace(comment)
		// empty string
		if comment == "" {
			comment = "-"
		}
		r = append(r, comment)
	}
	return strconv.Quote(strings.Join(r, ","))
}

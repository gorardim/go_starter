package {{.Package}}

import "github.com/gin-gonic/gin"

type {{cameCase .Name}} struct {
    svc {{.Name}}
}

{{range $v := .RpcArr}}
func (o *{{cameCase $.Name}}) {{$v.Name}}(c *gin.Context) (interface{}, error) {
	req := new({{$v.Request}})
	{{if needBindJson $v.Request}}
	if err := c.ShouldBind(req); err != nil {
        return nil, err
    }
    {{end}}
	resp, err := o.svc.{{$v.Name}}(c.Request.Context(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
{{end}}

func Register{{.Name}}(r *gin.Engine,svc {{.Name}},handle func(func(c *gin.Context) (interface{}, error))gin.HandlerFunc,middlewares ...gin.HandlerFunc) {
	server := &{{cameCase .Name}}{
        svc: svc,
    }
	{{range $v := .RpcArr}}
	r.{{method $v.Method}}("{{$v.Path}}", append(middlewares,handle(server.{{$v.Name}}))...){{end}}
}

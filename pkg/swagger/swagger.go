package swagger

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed dist
var swaggerUI embed.FS

func Router(router *gin.Engine, path string, swaggerContent string) {
	r := router.Group(path)

	r.GET("/_swagger/content/openapi.yaml", func(c *gin.Context) {
		c.String(http.StatusOK, swaggerContent)
	})

	r.StaticFS("/_swagger/static", http.FS(swaggerUI))

	tpl := template.Must(template.New("").ParseFS(swaggerUI, "dist/*.html"))
	router.SetHTMLTemplate(tpl)

	r.GET("/_swagger", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

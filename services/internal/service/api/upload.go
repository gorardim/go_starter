package api

import (
	"app/services/internal/component"

	"github.com/gin-gonic/gin"
)

type Upload struct {
	UploadComponent *component.UploadComponent
}

func (u *Upload) UploadImage(c *gin.Context) (interface{}, error) {
	return u.UploadComponent.UploadImage(c)
}

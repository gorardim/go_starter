package admin

import (
	"app/api/admin"
	"app/api/model"
	"app/pkg/errx"
	"app/services/internal/component"
	"app/services/internal/repo"

	"github.com/gin-gonic/gin"
)

type Upload struct {
	UploadComponent *component.UploadComponent
	*repo.AdminImageRepo
	*repo.AdminVideoRepo
}

func (u *Upload) UploadImage(c *gin.Context) (interface{}, error) {
	image, err := u.UploadComponent.UploadImage(c)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	adminImage := model.AdminImage{
		Url:  image.Url,
		Name: image.Name,
	}
	if err := u.AdminImageRepo.Create(c, &adminImage); err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	return &UploadImageResponse{
		Id:  adminImage.Id,
		Url: image.Url,
	}, nil
}

type UploadImageResponse struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func (u *Upload) UploadVideo(c *gin.Context) (interface{}, error) {
	file, err := u.UploadComponent.UploadVideo(c)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	adminVideo := model.AdminVideo{
		Url:  file.Url,
		Name: file.Name,
	}
	if err := u.AdminVideoRepo.Create(c, &adminVideo); err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	return &UploadVideoResponse{
		Id:  adminVideo.Id,
		Url: file.Url,
	}, nil
}

type UploadVideoResponse struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

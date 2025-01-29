package component

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"app/api/api"
	"app/pkg/errx"
	"app/pkg/randx"
	"app/services/internal/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/gin-gonic/gin"
)

type UploadComponent struct {
	Config *config.Config
	Client *oss.Client
}

type UploadImageResponse struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

func (u *UploadComponent) UploadImage(c *gin.Context) (*UploadImageResponse, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	now := time.Now()
	// get file ext
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("/images/%s/%s%s", now.Format("20060102"), randx.Seq(16), ext)
	dst := u.Config.UploadFileDir + filename
	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	if err = c.SaveUploadedFile(file, dst); err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	// save to oss
	bucket, err := u.Client.Bucket(u.Config.OssBucketName)
	if err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	if err = bucket.PutObjectFromFile(strings.TrimPrefix(filename, "/"), dst); err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	return &UploadImageResponse{
		Url:  u.Config.CdnUrl + filename,
		Name: file.Filename,
	}, nil
}

type UploadVideoResponse struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

func (u *UploadComponent) UploadVideo(c *gin.Context) (*UploadVideoResponse, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	now := time.Now()
	// get file ext
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("/videos/%s/%s%s", now.Format("20060102"), randx.Seq(16), ext)
	dst := u.Config.UploadFileDir + filename
	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	if err = c.SaveUploadedFile(file, dst); err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	bucket, err := u.Client.Bucket(u.Config.OssBucketName)
	if err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	if err = bucket.PutObjectFromFile(strings.TrimPrefix(filename, "/"), dst); err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	return &UploadVideoResponse{
		Url:  u.Config.CdnUrl + filename,
		Name: file.Filename,
	}, nil
}

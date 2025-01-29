package provider

import (
	"app/services/internal/config"
	"app/services/internal/pkg/avater"
	"net/http"
)

func NewAvaterClient(conf *config.Config, Client *http.Client) *avater.Client {
	return avater.NewClient(Client, conf.CdnUrl, conf.UploadFileDir)
}

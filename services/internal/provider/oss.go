package provider

import (
	"fmt"

	"app/services/internal/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewOssClient(config *config.Config) *oss.Client {
	client, err := oss.New(config.Oss.Endpoint, config.Oss.AccessKeyId, config.Oss.AccessKeySecret)
	if err != nil {
		panic(fmt.Errorf("new oss client error: %v", err))
	}
	return client
}

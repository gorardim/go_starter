package alioss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

func NewClient(endpoint, accessKeyId, accessKeySecret string) (*oss.Client, error) {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	return client, nil
}

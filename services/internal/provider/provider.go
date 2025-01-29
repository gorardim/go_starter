package provider

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewCaptcha,
	NewClient,
	NewNsqProducer,
	NewRedis,
	NewDb,
	NewAwsSesClient,
	NewAvaterClient,
	NewOssClient,
)

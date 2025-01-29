package provider

import (
	"app/services/internal/config"
	"app/services/internal/pkg/aws_ses"
)

func NewAwsSesClient(conf *config.Config) *aws_ses.Client {
	return aws_ses.NewClient(
		conf.Aws.Region,
		conf.Aws.AccessKeyId,
		conf.Aws.AccessSecret,
		conf.Aws.From,
	)
}

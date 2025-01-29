package conf

import (
	_ "embed"

	"app/pkg/gormx"
)

//go:embed app.local.yml
var AppLocalContent string

//go:embed app.dev.yml
var AppDevContent string

type AppConfig struct {
	AppName  string                   `yaml:"app_name"`
	Env      string                   `yaml:"env"`
	Database map[string]*gormx.Config `yaml:"database"`
	Redis    map[string]*RedisConfig  `yaml:"redis"`
	Log      struct {
		Mode string `yaml:"mode"`
		Dir  string `yaml:"dir"`
	} `yaml:"log"`
	UploadFileDir string `yaml:"upload_file_dir"`
	CdnUrl        string `yaml:"cdn_url"`

	Nsq struct {
		// 地址
		Addr string `json:"addr" yaml:"addr"`
	} `yaml:"nsq"`

	Jwt struct {
		OpenApiSecret string `yaml:"open_api_secret"`
	} `yaml:"jwt"`

	Aws struct {
		Region       string `yaml:"region"`
		AccessKeyId  string `yaml:"access_key_id"`
		AccessSecret string `yaml:"access_secret"`
		From         string `yaml:"from"`
	} `yaml:"aws"`
	Oss struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyId     string `yaml:"access_key_id"`
		AccessKeySecret string `yaml:"access_key_secret"`
	} `yaml:"oss"`
	OssBucketName string `yaml:"oss_bucket_name"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

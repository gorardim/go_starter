package config

import (
	"fmt"

	"app/component/conf"
	"gopkg.in/yaml.v3"
)

type Config struct {
	conf.AppConfig
	AppName string
	Env     string
}

func NewConfig(content string) *Config {
	c := &Config{}
	err := yaml.Unmarshal([]byte(content), &c.AppConfig)
	if err != nil {
		panic(fmt.Sprintf("unmarshal config err:%v", err))
	}
	return c
}

package provider

import (
	"log"

	"app/pkg/logger"
	"app/services/internal/config"
)

func InitLogger(conf *config.Config) {
	var l logger.Logger
	switch conf.Log.Mode {
	case "file":
		l = logger.NewFileLogger(conf.Log.Dir, conf.AppName)
	case "std":
		l = logger.NewStdLogger()
	default:
		log.Fatalf("unknown log mode: %s", conf.Log.Mode)
	}
	logger.DefaultLogger = l
}

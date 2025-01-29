package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"app/pkg/alert"
	"app/services/internal/component/lang"
	"app/services/internal/provider"

	"github.com/urfave/cli/v2"

	"app/component/conf"
	"app/pkg/cliutil"
	"app/services/internal/config"
)

func main() {
	restArg, options := cliutil.ParseOptions(os.Args, []string{"env", "app", "config"})
	env := options["env"]
	appName := options["app"]
	configFile := options["config"]
	if env == "" {
		env = "local"
	}
	if appName == "" {
		appName = "travel"
	}

	var c *config.Config
	if configFile != "" {
		file, err := os.ReadFile(configFile)
		if err != nil {
			log.Fatalf("failed to read config file: %v", err)
		}
		c = config.NewConfig(string(file))
	} else {
		switch env {
		case "local":
			c = config.NewConfig(conf.AppLocalContent)
		case "dev":
			c = config.NewConfig(conf.AppDevContent)
		default:
			log.Fatalf("unknown env: %s", env)
		}
	}
	alert.Env = env
	c.AppName = appName
	c.Env = env
	provider.InitLogger(c)

	if err := lang.Init(); err != nil {
		alert.Alert(context.Background(), "lang init failed", []string{
			fmt.Sprintf("err: %v", err),
		})
	}
	app := &cli.App{
		Commands: newCommand(c),
	}
	if err := app.Run(restArg); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}

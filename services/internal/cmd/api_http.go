package cmd

import (
	"app/services/internal/router"

	"github.com/urfave/cli/v2"
)

type ApiHttpCmd *cli.Command

func NewApiHttpCmd(router *router.ApiRouter) ApiHttpCmd {
	return &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "addr",
				Usage: "listen address",
				Value: ":7702",
			},
		},
		Name:  "api",
		Usage: "api",
		Action: func(c *cli.Context) error {
			addr := c.String("addr")
			return router.Run(addr)
		},
	}
}

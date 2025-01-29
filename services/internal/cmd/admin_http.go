package cmd

import (
	"app/services/internal/router"

	"github.com/urfave/cli/v2"
)

type AdminHttpCmd *cli.Command

func NewAdminHttpCmd(router *router.AdminRouter) AdminHttpCmd {
	return &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "addr",
				Usage: "listen address",
				Value: ":7703",
			},
		},
		Name:  "admin",
		Usage: "admin",
		Action: func(c *cli.Context) error {
			addr := c.String("addr")
			return router.Run(addr)
		},
	}
}

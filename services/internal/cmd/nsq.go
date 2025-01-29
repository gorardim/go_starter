package cmd

import (
	"app/api/job"
	"app/pkg/broker/nsqx"
	"app/pkg/cliutil"
	"app/services/internal/config"

	"github.com/urfave/cli/v2"
)

type NsqCmd *cli.Command

func NewNsqCmd(
	conf *config.Config,
	userRegisterConsumer job.UserRegisterConsumer,
) NsqCmd {
	return &cli.Command{
		Name:  "nsq",
		Usage: "nsq",
		Action: func(cliCtx *cli.Context) error {
			_, shutdown := cliutil.NewShutDown()
			job.UserRegisterDefaultConsumer(userRegisterConsumer, conf.Nsq.Addr, nsqx.MaxInFlight(10))
			shutdown()
			return nil
		},
	}
}

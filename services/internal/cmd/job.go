package cmd

import (
	"context"
	"fmt"
	"time"

	"app/pkg/alert"
	"app/pkg/cronx"

	"github.com/urfave/cli/v2"
)

type JobCmd *cli.Command

func NewJobCmd() JobCmd {
	return &cli.Command{
		Name:  "job",
		Usage: "job",
		Action: func(cliCtx *cli.Context) error {
			cron := cronx.NewCron()
			// 每天0点执行
			cron.AddFunc("收益释放", "15 0 * * *", func(ctx context.Context) error {
				alert.Alert(ctx, "all job done", []string{
					fmt.Sprintf("time: %s", time.Now().Format(time.DateTime)),
				})
				return nil
			})
			cron.Run()
			return nil
		},
	}
}

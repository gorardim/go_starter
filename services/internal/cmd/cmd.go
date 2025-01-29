package cmd

import (
	"app/pkg/cliutil"

	"github.com/urfave/cli/v2"
)

type Commands struct {
	ApiHttpCmd   ApiHttpCmd
	AdminHttpCmd AdminHttpCmd
	NsqCmd       NsqCmd
	JobCmd       JobCmd
}

func NewCliCommands(commands *Commands) []*cli.Command {
	return cliutil.NewCliCommand(commands)
}

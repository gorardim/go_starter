package cmd

import "github.com/google/wire"

var Provider = wire.NewSet(
	wire.Struct(new(Commands), "*"),
	NewCliCommands,
	NewApiHttpCmd,
	NewAdminHttpCmd,
	NewNsqCmd,
	NewJobCmd,
)

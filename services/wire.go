//go:build wireinject
// +build wireinject

package main

import (
	apijob "app/api/job"
	"app/component"
	"app/services/internal/cache"
	"app/services/internal/cmd"
	appcomponent "app/services/internal/component"
	"app/services/internal/config"
	"app/services/internal/provider"
	"app/services/internal/repo"
	"app/services/internal/router"
	"app/services/internal/service/admin"
	"app/services/internal/service/api"
	"app/services/internal/service/job"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

var providerSet = wire.NewSet(
	cmd.Provider,
	api.Provider,
	router.Provider,
	repo.Provider,
	provider.Provider,
	component.Provider,
	job.Provider,
	apijob.Provider,
	appcomponent.Provider,
	admin.Provider,
	cache.Provider,
)

func newCommand(conf *config.Config) []*cli.Command {
	panic(wire.Build(providerSet))
}

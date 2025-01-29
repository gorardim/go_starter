package api

import (
	"app/api/api"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	wire.Struct(new(Auth), "*"),
	wire.Bind(new(api.AuthServer), new(*Auth)),
	wire.Struct(new(Upload), "*"),
)

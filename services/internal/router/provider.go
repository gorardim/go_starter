package router

import "github.com/google/wire"

var Provider = wire.NewSet(
	wire.Struct(new(ApiRouter), "*"),
	wire.Struct(new(AdminRouter), "*"),
)

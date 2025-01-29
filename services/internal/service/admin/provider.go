package admin

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	wire.Struct(new(Auth), "*"),
	wire.Struct(new(CenterMenu), "*"),
	wire.Struct(new(CenterUser), "*"),
	wire.Struct(new(CenterRole), "*"),
	wire.Struct(new(CenterApi), "*"),
	wire.Struct(new(Upload), "*"),
	wire.Struct(new(AdminImage), "*"),
	wire.Struct(new(AdminVideo), "*"),
	wire.Struct(new(User), "*"),
	wire.Struct(new(Setting), "*"),
)

package cache

import "github.com/google/wire"

var Provider = wire.NewSet(
	wire.Struct(new(Home), "*"),
	wire.Struct(new(OneKeyReceiveIncomeLock), "*"),
	wire.Struct(new(UserResetPasswordLimiter), "*"),
	wire.Struct(new(PayPasswordLimiter), "*"),
)

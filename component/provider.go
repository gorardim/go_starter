package component

import (
	"app/component/counter"
	"app/component/locks"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	wire.Struct(new(counter.Counter), "*"),
	wire.Struct(new(locks.RedisLock), "*"),
)

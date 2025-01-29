package job

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewUserRegisterPublisher,
)

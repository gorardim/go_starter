package job

import (
	"app/api/job"
	"app/services/internal/service/job/userlevel"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	wire.Struct(new(userlevel.UserRegister), "*"),
	wire.Bind(new(job.UserRegisterConsumer), new(*userlevel.UserRegister)),
)

package component

import (
	"app/services/internal/component/email"
	setting "app/services/internal/component/settting"
	"app/services/internal/component/twofa"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	setting.Provider,
	wire.Struct(new(UploadComponent), "*"),
	wire.Struct(new(AdminPermissionComponent), "*"),
	wire.Struct(new(UserAuthComponent), "*"),
	wire.Struct(new(email.Email), "*"),
	wire.Struct(new(twofa.TwoFA), "*"),
)

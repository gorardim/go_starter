//go:build wireinject
// +build wireinject

package setting

import (
	"app/services/internal/repo"

	"github.com/google/wire"
)

func NewSetting(repo *repo.SettingRepo) *Setting {
	panic(wire.Build(Provider))
}

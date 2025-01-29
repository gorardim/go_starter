package setting

import (
	"app/api/model"

	"github.com/google/wire"
)

type IntConfig = setting[int]
type FloatConfig = setting[float64]
type FloatListConfig = setting[[]float64]
type StringConfig = setting[string]
type StringListConfig = setting[[]string]
type UsdtNetworkListConfig = setting[[]*UsdtNetwork]
type LangTypeConfig = setting[model.LangType]
type CurrencExchangeRateList = setting[[]*CurrencyExchangeRateItem]
type DepositCurrencyExchangeRateList = setting[[]*DepositCurrencyExchangeRateItem]
type ConfigFixedTopDiscoveryHomePage = setting[[]*model.FixedTopDiscoveryHomePage]

type Setting struct {
	IntConfig                       *IntConfig
	FloatConfig                     *FloatConfig
	StringConfig                    *StringConfig
	FloatListConfig                 *FloatListConfig
	StringListConfig                *StringListConfig
	UsdtNetworkListConfig           *UsdtNetworkListConfig
	LangTypeConfig                  *LangTypeConfig
	CurrencExchangeRate             *CurrencExchangeRateList
	DepositCurrencyExchangeRateList *DepositCurrencyExchangeRateList
	ConfigFixedTopDiscoveryHomePage *ConfigFixedTopDiscoveryHomePage
}

var Provider = wire.NewSet(
	wire.Struct(new(Setting), "*"),
	wire.Struct(new(IntConfig), "*"),
	wire.Struct(new(FloatConfig), "*"),
	wire.Struct(new(StringConfig), "*"),
	wire.Struct(new(FloatListConfig), "*"),
	wire.Struct(new(StringListConfig), "*"),
	wire.Struct(new(UsdtNetworkListConfig), "*"),
	wire.Struct(new(LangTypeConfig), "*"),
	wire.Struct(new(CurrencExchangeRateList), "*"),
	wire.Struct(new(DepositCurrencyExchangeRateList), "*"),
	wire.Struct(new(ConfigFixedTopDiscoveryHomePage), "*"),
)

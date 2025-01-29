package admin

import (
	"app/api/model"
	"context"
)

type SettingServer interface {
	//x:api post /admin/setting/deposit_cash_rate_country_config 充值订单国家配置
	DepositCashRateCountryConfig(ctx context.Context, req *DepositCashRateCountryConfigRequest) (*DepositCashRateCountryConfigResponse, error)
	//x:api post /admin/setting/deposit_cash_rate_country_config_update 更新充值订单国家配置
	DepositCashRateCountryConfigUpdate(ctx context.Context, req *DepositCashRateCountryConfigUpdateRequest) (*DepositCashRateCountryConfigUpdateResponse, error)
	//x:api post /admin/setting/deposit_cash_rate_country_config_push 推送充值订单国家配置
	DepositCashRateCountryConfigPush(ctx context.Context, req *DepositCashRateCountryConfigPushRequest) (*DepositCashRateCountryConfigPushResponse, error)
}

type DepositCashRateCountryConfigPushRequest struct {
	// deposit currency exchange rate item
	DepositCurrencyExchangeRateItem *DepositCurrencyExchangeRateItem `json:"deposit_currency_exchange_rate_item"`
}

type DepositCashRateCountryConfigPushResponse struct {
}

type DepositCashRateCountryConfigUpdateRequest struct {
	// deposit currency exchange rate list
	DepositCurrencyExchangeRateList []*DepositCurrencyExchangeRateItem `json:"deposit_currency_exchange_rate_list"`
}

type DepositCashRateCountryConfigUpdateResponse struct {
}

type DepositCashRateCountryConfigRequest struct {
}

type DepositCashRateCountryConfigResponse struct {
	// deposit currency exchange rate list
	DepositCurrencyExchangeRateList []*DepositCurrencyExchangeRateItem `json:"deposit_currency_exchange_rate_list"`
}

type DepositCurrencyExchangeRateItem struct {
	// currency code
	CurrencyCode string `json:"currency_code"`
	// Rate
	Rate string `json:"rate"`
	// symbol
	Symbol string `json:"symbol"`
	// bank card
	BankCard string `json:"bank_card"`
	// country name
	CountryName model.Lang `json:"country_name"`
	// country code: USA, CHN, ...
	CountryCode string `json:"country_code"`
	//
	Rules []*model.DisplaySettingItem `json:"rules"`
	//
	RuleDesc string `json:"rule_desc"`
}

package admin

import (
	"app/api/admin"
	"app/api/api"
	"app/api/model"
	"app/pkg/errx"
	"app/pkg/gormx"
	setting "app/services/internal/component/settting"
	"context"
	"fmt"
	"strconv"
)

var _ admin.SettingServer = (*Setting)(nil)

type Setting struct {
	Setting *setting.Setting
}

func (s *Setting) DepositCashRateCountryConfig(ctx context.Context, req *admin.DepositCashRateCountryConfigRequest) (*admin.DepositCashRateCountryConfigResponse, error) {
	rateList, err := s.Setting.DepositCurrencyExchangeRateList.GetValue(ctx, setting.DepositCurrencyExchangeRate)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err.Error())
	}

	depositCurrencyExchangeRateList := make([]*admin.DepositCurrencyExchangeRateItem, 0, len(rateList))
	for _, item := range rateList {
		depositCurrencyExchangeRateList = append(depositCurrencyExchangeRateList, &admin.DepositCurrencyExchangeRateItem{
			CurrencyCode: item.Name,
			Rate:         fmt.Sprintf("%.4f", item.Rate),
			Symbol:       item.Symbol,
			BankCard:     item.BankCard,
			CountryName:  item.CountryName.Data,
			CountryCode:  item.CountryCode,
			Rules:        item.Rules,
			RuleDesc:     item.RuleDesc,
		})
	}

	return &admin.DepositCashRateCountryConfigResponse{
		DepositCurrencyExchangeRateList: depositCurrencyExchangeRateList,
	}, nil
}

func (s *Setting) DepositCashRateCountryConfigUpdate(ctx context.Context, req *admin.DepositCashRateCountryConfigUpdateRequest) (*admin.DepositCashRateCountryConfigUpdateResponse, error) {

	depositRateCountry := make([]*setting.DepositCurrencyExchangeRateItem, 0)
	for _, v := range req.DepositCurrencyExchangeRateList {
		rate, err := strconv.ParseFloat(v.Rate, 64)
		if err != nil {
			return nil, errx.New(admin.ErrInvalidParam, "收益率格式错误")
		}
		depositRateCountry = append(depositRateCountry, &setting.DepositCurrencyExchangeRateItem{

			Name:        v.CurrencyCode,
			Rate:        rate,
			Symbol:      v.Symbol,
			BankCard:    v.BankCard,
			CountryCode: v.CountryCode,
			CountryName: gormx.JsonObject[model.Lang]{Data: v.CountryName},
			Rules:       v.Rules,
			RuleDesc:    v.RuleDesc,
		})
	}

	if err := s.Setting.DepositCurrencyExchangeRateList.SetValue(ctx, setting.DepositCurrencyExchangeRate, depositRateCountry); err != nil {
		return nil, errx.New(admin.ErrBusiness, err.Error())
	}
	return &admin.DepositCashRateCountryConfigUpdateResponse{}, nil
}

func (s *Setting) DepositCashRateCountryConfigPush(ctx context.Context, req *admin.DepositCashRateCountryConfigPushRequest) (*admin.DepositCashRateCountryConfigPushResponse, error) {
	rateList, err := s.Setting.DepositCurrencyExchangeRateList.GetValue(ctx, setting.DepositCurrencyExchangeRate)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err.Error())
	}

	rate, err := strconv.ParseFloat(req.DepositCurrencyExchangeRateItem.Rate, 64)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, "收益率格式错误")
	}

	rateList = append(rateList, &setting.DepositCurrencyExchangeRateItem{

		Name:        req.DepositCurrencyExchangeRateItem.CurrencyCode,
		Rate:        rate,
		Symbol:      req.DepositCurrencyExchangeRateItem.Symbol,
		BankCard:    req.DepositCurrencyExchangeRateItem.BankCard,
		CountryCode: req.DepositCurrencyExchangeRateItem.CountryCode,
		CountryName: gormx.JsonObject[model.Lang]{Data: req.DepositCurrencyExchangeRateItem.CountryName},
		Rules:       req.DepositCurrencyExchangeRateItem.Rules,
		RuleDesc:    req.DepositCurrencyExchangeRateItem.RuleDesc,
	})

	if err := s.Setting.DepositCurrencyExchangeRateList.SetValue(ctx, setting.DepositCurrencyExchangeRate, rateList); err != nil {
		return nil, errx.New(admin.ErrBusiness, err.Error())
	}

	return &admin.DepositCashRateCountryConfigPushResponse{}, nil
}

// Code generated by xapi. DO NOT EDIT.

package admin

import "github.com/gin-gonic/gin"

type settingServer struct {
	svc SettingServer
}

func (o *settingServer) DepositCashRateCountryConfig(c *gin.Context) (interface{}, error) {
	req := new(DepositCashRateCountryConfigRequest)

	resp, err := o.svc.DepositCashRateCountryConfig(c.Request.Context(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *settingServer) DepositCashRateCountryConfigUpdate(c *gin.Context) (interface{}, error) {
	req := new(DepositCashRateCountryConfigUpdateRequest)

	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	resp, err := o.svc.DepositCashRateCountryConfigUpdate(c.Request.Context(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *settingServer) DepositCashRateCountryConfigPush(c *gin.Context) (interface{}, error) {
	req := new(DepositCashRateCountryConfigPushRequest)

	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	resp, err := o.svc.DepositCashRateCountryConfigPush(c.Request.Context(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func RegisterSettingServer(r *gin.Engine, svc SettingServer, handle func(func(c *gin.Context) (interface{}, error)) gin.HandlerFunc, middlewares ...gin.HandlerFunc) {
	server := &settingServer{
		svc: svc,
	}

	r.POST("/admin/setting/deposit_cash_rate_country_config", append(middlewares, handle(server.DepositCashRateCountryConfig))...)
	r.POST("/admin/setting/deposit_cash_rate_country_config_update", append(middlewares, handle(server.DepositCashRateCountryConfigUpdate))...)
	r.POST("/admin/setting/deposit_cash_rate_country_config_push", append(middlewares, handle(server.DepositCashRateCountryConfigPush))...)
}

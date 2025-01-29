package bsc

import "context"

// 查询余额 /get_usdt_balance

type GetUsdtBalanceRequest struct {
	Appid string `json:"appid" url:"appid"` // 商户ID
}

type GetUsdtBalanceResponse struct {
	Appid   string `json:"appid"`   // AppID
	Balance string `json:"balance"` // 余额
}

func (c *Client) GetUsdtBalance(ctx context.Context) (*Response[GetUsdtBalanceResponse], error) {
	resp := &Response[GetUsdtBalanceResponse]{}
	err := c.postForm(ctx, "/get_usdt_balance", &GetUsdtBalanceRequest{
		Appid: c.appid,
	}, resp, WithSign())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

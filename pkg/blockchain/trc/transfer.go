package trc

import "context"

// 发起转账（提币）*此接口要加ip白名单才可以请求 /transfer

type TransferRequest struct {
	Appid    string `json:"appid" url:"appid"`         // 商户ID
	Oid      string `json:"oid" url:"oid"`             // 商户提币订单号
	CoinCode string `json:"coin_code" url:"coin_code"` // 币种 USDT
	Address  string `json:"address" url:"address"`     // 提币地址
	Amount   string `json:"amount" url:"amount"`       // 提币金额
}

type TransferResponse struct {
	Appid    string `json:"appid"`      // AppID
	Uid      string `json:"uid"`        // 传参的UID
	ChanType string `json:"chain_type"` // tron
	Address  string `json:"address"`    // 用户地址
	Sign     string `json:"sign"`       // 签名
}

func (c *Client) Transfer(ctx context.Context, req *TransferRequest) (*Response[TransferResponse], error) {
	resp := &Response[TransferResponse]{}
	err := c.postForm(ctx, "/transfer", req, resp, WithSign())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

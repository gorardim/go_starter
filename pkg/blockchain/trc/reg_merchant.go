package trc

import "context"

type RegMerchantRequest struct {
	Name              string `json:"name" url:"name"`                               // 商户名
	Email             string `json:"email" url:"email"`                             // 邮箱
	Mobile            string `json:"mobile" url:"mobile"`                           // 手机
	Pwd               string `json:"pwd" url:"pwd"`                                 // 密码
	RechargeNotifyUrl string `json:"recharge_notify_url" url:"recharge_notify_url"` // 充值回调
	TransferNotifyUrl string `json:"transfer_notify_url" url:"transfer_notify_url"` // 提币回调
}

type RegMerchantResponse struct {
	Appid      string `json:"appid"` // AppID
	Name       string `json:"name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	HotAddress string `json:"hot_address"`  // 热钱包地址
	TwoAuthKey string `json:"two_auth_key"` // 两步验证秘钥
	SecretKey  string `json:"secret_key"`   // 秘钥
	Sign       string `json:"sign"`         // 签名
}

// RegMerchant 注册商户
func (c *Client) RegMerchant(ctx context.Context, req *RegMerchantRequest) (*Response[RegMerchantResponse], error) {
	resp := &Response[RegMerchantResponse]{}
	err := c.postForm(ctx, "/reg_merchant", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

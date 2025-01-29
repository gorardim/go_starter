package trc

import "context"

// 获取用户充币地址 /get_address

type getAddressRequest struct {
	Appid     string `json:"appid" url:"appid"`           // 商户ID
	Uid       string `json:"uid" url:"uid"`               // 前缀+用户唯一ID
	ChainType string `json:"chain_type" url:"chain_type"` // tron
}

type GetAddressResponse struct {
	Appid     string `json:"appid"`      // AppID
	Uid       string `json:"uid"`        // 传参的UID
	ChainType string `json:"chain_type"` // tron
	Address   string `json:"address"`    // 用户地址
	Sign      string `json:"sign"`       // 签名
}

func (c *Client) GetAddress(ctx context.Context, uid string) (*Response[GetAddressResponse], error) {
	req := &getAddressRequest{
		Appid:     c.appid,
		Uid:       uid,
		ChainType: "tron",
	}
	resp := &Response[GetAddressResponse]{}
	err := c.postForm(ctx, "/get_address", req, resp, WithSign())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

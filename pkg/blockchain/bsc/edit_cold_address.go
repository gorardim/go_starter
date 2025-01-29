package bsc

import "context"

type EditColdAddressRequest struct {
	AppId       string `json:"appid" url:"appid"`
	TwoAuthCode string `json:"two_auth_code" url:"two_auth_code"`
	ChainType   string `json:"chain_type" url:"chain_type"`
	Address     string `json:"address" url:"address"`
}

type EditColdAddressResponse struct {
}

func (c *Client) EditColdAddress(ctx context.Context, address, twoAuthCode string) (resp *Response[EditColdAddressResponse], err error) {
	resp = &Response[EditColdAddressResponse]{}
	req := &EditColdAddressRequest{
		AppId:       c.appid,
		TwoAuthCode: twoAuthCode,
		ChainType:   "bsc",
		Address:     address,
	}
	if err = c.postForm(ctx, "/edit_cold_address", req, resp, WithSign()); err != nil {
		return nil, err
	}
	return resp, nil
}

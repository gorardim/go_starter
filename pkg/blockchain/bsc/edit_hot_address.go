package bsc

import "context"

type EditHotAddressRequest struct {
	AppId       string `json:"appid" url:"appid"`
	ChainType   string `json:"chain_type" url:"chain_type"`
	TwoAuthCode string `json:"two_auth_code" url:"two_auth_code"`
}

type EditHotAddressResponse struct{}

func (c *Client) EditHotAddress(ctx context.Context, twoAuthCode string) (resp *Response[EditHotAddressResponse], err error) {
	resp = &Response[EditHotAddressResponse]{}
	req := &EditHotAddressRequest{
		AppId:       c.appid,
		ChainType:   "bsc",
		TwoAuthCode: twoAuthCode,
	}
	if err = c.postForm(ctx, "/change_hot_address", req, resp, WithSign()); err != nil {
		return nil, err
	}
	return resp, nil
}

package bsc

import "app/pkg/transport"

type config struct {
	baseUrl    string
	appid      string
	name       string
	hotAddress string
	twoAuthKey string
	secretKey  string
}

var testConfig = config{}

func newClient() *Client {
	return &Client{
		httpClient: transport.NewClient(&transport.Debug{}),
		baseUrl:    testConfig.baseUrl,
		secret:     testConfig.secretKey,
		appid:      testConfig.appid,
	}
}

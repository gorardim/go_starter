package avater

import "net/http"

type Client struct {
	HttpClient *http.Client
	Url        string
	d          string
	CdnUrl     string
	Storage    string
}

func NewClient(httpClient *http.Client, CdnUrl string, storage string) *Client {
	return &Client{
		HttpClient: httpClient,
		Url:        "https://www.gravatar.com/avatar/",
		CdnUrl:     CdnUrl,
		Storage:    storage,
	}
}

package bsc

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"sort"

	"app/pkg/hashutil"
	"github.com/google/go-querystring/query"
)

type Client struct {
	httpClient *http.Client
	baseUrl    string
	secret     string
	appid      string
}

func NewClient(httpClient *http.Client, baseUrl, secret string, appid string) *Client {
	return &Client{
		httpClient: httpClient,
		baseUrl:    baseUrl,
		secret:     secret,
		appid:      appid,
	}
}

func (c *Client) postForm(ctx context.Context, path string, req interface{}, res interface{}, ops ...Option) error {
	o := &options{}
	for _, op := range ops {
		op(o)
	}
	values, err := query.Values(req)
	if err != nil {
		return err
	}
	if o.sign {
		values.Set("sign", sign(values, c.secret))
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseUrl+path, bytes.NewBufferString(values.Encode()))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(all, res)
}

func sign(data url.Values, secret string) string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var signPars string
	for _, k := range keys {
		if k != "sign" {
			signPars += k + "=" + data.Get(k) + "&"
		}
	}
	signPars += "key=" + secret
	return hashutil.Md5(signPars)
}

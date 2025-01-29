package bsc

import "net/url"

// 充币回调

type RechargeNotifyRequest struct {
	Address  string `json:"address" url:"address" form:"address"`
	Amount   string `json:"amount" url:"amount" form:"amount"`
	Appid    string `json:"appid" url:"appid" form:"appid"`
	CoinCode string `json:"coin_code" url:"coin_code" form:"coin_code"`
	Sign     string `json:"sign" url:"sign" form:"sign"`
	TxId     string `json:"tx_id" url:"tx_id" form:"tx_id"`
	Uid      string `json:"uid" url:"uid" form:"uid"`
}

// 提币回调

type TransferNotifyRequest struct {
	Address  string `json:"address" url:"address" form:"address"`
	Amount   string `json:"amount" url:"amount" form:"amount"`
	Appid    string `json:"appid" url:"appid" form:"appid"`
	BillCode string `json:"bill_code" url:"bill_code" form:"bill_code"`
	CoinCode string `json:"coin_code" url:"coin_code" form:"coin_code"`
	Oid      string `json:"oid" url:"oid" form:"oid"`
	Sign     string `json:"sign" url:"sign" form:"sign"`
	Status   string `json:"status" url:"status" form:"status"`
	TxId     string `json:"tx_id" url:"tx_id" form:"tx_id"`
}

func CheckSign(appId, secret string, values url.Values) bool {
	if values.Get("appid") != appId {
		return false
	}
	s := values.Get("sign")
	if s == "" {
		return false
	}
	v := sign(values, secret)
	return s == v
}

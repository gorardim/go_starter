package trc

import "net/url"

// 充币回调

// address=TDZGNDRwWBWLALKzUdrZFE78Wtu8NHcYkR&
// amount=20&
// appid=287b63d939d23dee3e536e255d63b5a4&
// coin_code=usdt&
// sign=2a6b8ffc27b78cb54e1b3b2b26677999&
// tx_id=1b89d3cebb7cf857db4a1b18df22d67f12b3d79780ebd9f3d52b2577beb94259&uid
// =travel_test_trc1
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

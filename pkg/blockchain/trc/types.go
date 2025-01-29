package trc

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *T     `json:"data"`
}

const ChainTypeTron = "tron" //  usdt
const CoinCodeUsdt = "usdt"  // usdt

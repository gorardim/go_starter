package trc

import (
	"net/url"
	"testing"
)

func TestCheckSign(t *testing.T) {
	// address=THdxBQ6nmG2ZmpmwkAnuaGFzBzs9jCVZpa&amount=1&appid=178efcf524d0f4bf086577aa51fa10b8&coin_code=usdt&sign=bafa866da8aa94d333b8b8ba0c076eca&tx_id=994d1ea724d71fe80d1a629214a7e4a4bc23afee2992d19118979802d11e8507&uid=cixtest10001
	// curl -d 'address=THdxBQ6nmG2ZmpmwkAnuaGFzBzs9jCVZpa&amount=1&appid=178efcf524d0f4bf086577aa51fa10b8&coin_code=usdt&sign=bafa866da8aa94d333b8b8ba0c076eca&tx_id=994d1ea724d71fe80d1a629214a7e4a4bc23afee2992d19118979802d11e8507&uid=cixtest10001' http://127.0.0.1:7802/api/encrypay/recharge/notify
	// curl -d 'address=THdxBQ6nmG2ZmpmwkAnuaGFzBzs9jCVZpa&amount=1&appid=178efcf524d0f4bf086577aa51fa10b8&coin_code=usdt&sign=bafa866da8aa94d333b8b8ba0c076eca&tx_id=994d1ea724d71fe80d1a629214a7e4a4bc23afee2992d19118979802d11e8507&uid=cixtest10001' http://u-test.967s.com/api/encrypay/recharge/notify

	// address=TWRWfJCtpxQKVQNWVX63t5Ww81XNb1Lxpf&amount=1&appid=178efcf524d0f4bf086577aa51fa10b8&bill_code=518a16a0d59b1182b0315fa148cd9281&coin_code=usdt&oid=23031623243927223694&sign=0fa110a1747c56c3e500c4489f9f2256&status=2&tx_id=4adcfe0339f73fd42e33db342d6b07e6a4984f7e8e64177dacb7210bcc48cf9b

	values := url.Values{}
	values.Set("address", "THdxBQ6nmG2ZmpmwkAnuaGFzBzs9jCVZpa")
	values.Set("amount", "1")
	values.Set("appid", "178efcf524d0f4bf086577aa51fa10b8")
	values.Set("coin_code", "usdt")
	values.Set("sign", "bafa866da8aa94d333b8b8ba0c076eca")
	values.Set("tx_id", "994d1ea724d71fe80d1a629214a7e4a4bc23afee2992d19118979802d11e8507")
	values.Set("uid", "cixtest10001")

	// bafa866da8aa94d333b8b8ba0c076eca
	s := sign(values, "c612cf53ca73103d7a466f5af9874224")
	t.Log(s)
}

// 537bbda57d11267bc6de31f4411e95dc36354e0c0becfeb2c5f06326213db118
func TestCheckSign1(t *testing.T) {
	// address=TZ1WKehrZ7UDDKDXtBA6q7APEoQm7GZCXr&amount=9&appid=178efcf524d0f4bf086577aa51fa10b8&bill_code=d07043b9e75af69ede54fbc585016690&coin_code=usdt&oid=23031702415821115153&sign=92043594bf864cd43b14b4f94af78f62&status=2&tx_id=537bbda57d11267bc6de31f4411e95dc36354e0c0becfeb2c5f06326213db118
	// curl -d 'address=TZ1WKehrZ7UDDKDXtBA6q7APEoQm7GZCXr&amount=9&appid=178efcf524d0f4bf086577aa51fa10b8&bill_code=d07043b9e75af69ede54fbc585016690&coin_code=usdt&oid=23031702415821115153&sign=92043594bf864cd43b14b4f94af78f62&status=2&tx_id=537bbda57d11267bc6de31f4411e95dc36354e0c0becfeb2c5f06326213db118' http://u-test.967s.com/api/encrypay/transfer/notify
	values := url.Values{}
	values.Set("address", "TZ1WKehrZ7UDDKDXtBA6q7APEoQm7GZCXr")
	values.Set("amount", "9")
	values.Set("appid", "178efcf524d0f4bf086577aa51fa10b8")
	values.Set("bill_code", "d07043b9e75af69ede54fbc585016690")
	values.Set("coin_code", "usdt")
	values.Set("oid", "23031702415821115153")
	values.Set("sign", "92043594bf864cd43b14b4f94af78f62")
	values.Set("status", "2")
	values.Set("tx_id", "537bbda57d11267bc6de31f4411e95dc36354e0c0becfeb2c5f06326213db118")

	// 92043594bf864cd43b14b4f94af78f62
	s := sign(values, "c612cf53ca73103d7a466f5af9874224")
	t.Log(s)
}

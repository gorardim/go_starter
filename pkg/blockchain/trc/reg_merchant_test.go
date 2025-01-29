package trc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_RegMerchant(t *testing.T) {
	merchant, err := newClient().RegMerchant(context.Background(), &RegMerchantRequest{
		Name:              "cix test",
		Email:             "cix@cix",
		Mobile:            "13100050000",
		Pwd:               "xKqnoAhgUr6JPf42mlB54Sa1MI0YBdQ5",
		RechargeNotifyUrl: "http://u.967s.com/encrypay/recharge/notify",
		TransferNotifyUrl: "http://u.967s.com/encrypay/transfer/notify",
	})
	assert.NoError(t, err)
	t.Log(merchant)
}

func TestClient_RegMerchant1(t *testing.T) {
	merchant, err := newClient().RegMerchant(context.Background(), &RegMerchantRequest{
		Name:              "cix",
		Email:             "wangpengdylan@gmail.com",
		Mobile:            "855011858180",
		Pwd:               "Dy#lan688@1z",
		RechargeNotifyUrl: "http://u.967s.com/api/encrypay/recharge/notify",
		TransferNotifyUrl: "http://u.967s.com/api/encrypay/transfer/notify",
	})
	assert.NoError(t, err)
	t.Log(merchant)
}

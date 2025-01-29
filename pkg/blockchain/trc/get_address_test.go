package trc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetAddress(t *testing.T) {
	address, err := newClient().GetAddress(context.Background(), "traveltesta2")
	assert.NoError(t, err)
	t.Log(address)
	// {"code":200,"msg":"成功","data":{"appid":"287b63d939d23dee3e536e255d63b5a4","uid":"traveltesta1","chain_type":"tron","address":"TXEgiTXscBaS5JQczGMfinnnpbEWD3Bh4V","sign":"6770587c97714bd2c9296d0e1658217e"}}
}

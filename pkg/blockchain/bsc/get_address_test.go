package bsc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetAddress(t *testing.T) {
	address, err := newClient().GetAddress(context.Background(), "traveltesta1")
	assert.NoError(t, err)
	t.Log(address)
	// {"code":200,"msg":"成功","data":{"appid":"92a13b10ee5e5882b9cafeb393546751","uid":"traveltesta1","address":"0x0D02f35Ce2207c5A6783206F729f63330611f03B","sign":"7b4ff26c0b812b95b1d445cf8a6cda47"}}

}

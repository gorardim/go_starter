package randx

import "testing"

func TestGenUniqueId(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GenUniqueId())
	}
}

func TestVerifyUniqueId(t *testing.T) {
	tradeNo := GenUniqueId()
	if got := VerifyUniqueId(tradeNo); !got {
		t.Errorf("VerifyUniqueId() = %v", got)
	}
}

package userwallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encrypt(t *testing.T) {
	bytes, err := encrypt([]byte("123456"), []byte("0123456789abcdef0123456789abcdef"))
	assert.NoError(t, err)
	t.Log(bytes)
	d, err := decrypt(bytes, []byte("0123456789abcdef0123456789abcdef"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("123456"), d)
}

func TestDecryptUserAddress(t *testing.T) {
	address, err := EncryptUserAddress("0x55d398326f99059ff775485246999027b3197955", "0123456789abcdef0123456789abcdef")
	assert.NoError(t, err)
	t.Log(address)
	address, err = DecryptUserAddress(address, "0123456789abcdef0123456789abcdef")
	assert.NoError(t, err)
	assert.Equal(t, "0x55d398326f99059ff775485246999027b3197955", address)
}

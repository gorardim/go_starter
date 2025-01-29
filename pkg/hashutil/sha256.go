package hashutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSha256(data []byte, secret string) (string, error) {
	hash := hmac.New(sha256.New, []byte(secret))
	_, err := hash.Write(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

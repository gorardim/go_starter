package randx

import (
	"math/rand"
	"time"
)

const (
	alphaNum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	digit    = "0123456789"
	alpha    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Int(number int) int {
	randInt := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randInt.Intn(number)
}

// Seq 生成随机序列
func Seq(length int) string {
	var r = make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = alphaNum[Int(len(alphaNum))]
	}
	return string(r)
}

// Digit 生成数字序列
func Digit(length int) string {
	var r = make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = digit[Int(len(digit))]
	}
	return string(r)
}

func Alpha(length int) string {
	var r = make([]byte, length)
	for i := 0; i < length; i++ {
		r[i] = alpha[Int(len(alpha))]
	}
	return string(r)
}

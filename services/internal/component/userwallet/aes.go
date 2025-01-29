package userwallet

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func PKCS5Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func PKCS5UnPadding(plainText []byte, blockSize int) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number >= length || number > blockSize {
		return nil, fmt.Errorf("padding size error")
	}
	return plainText[:length-number], nil
}

func encrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 填充原文
	blockSize := block.BlockSize()
	rawData = PKCS5Padding(rawData, blockSize)
	// 初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, blockSize+len(rawData))
	// block大小 16
	iv := cipherText[:blockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	// block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func decrypt(encryptData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := encryptData[:blockSize]
	encryptData = encryptData[blockSize:]

	// CBC mode always works in whole blocks.
	if len(encryptData)%blockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(encryptData, encryptData)
	// 解填充
	return PKCS5UnPadding(encryptData, blockSize)
}

func DecryptUserAddress(encryptData, key string) (string, error) {
	// base64解码
	encryptDataByte, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return "", err
	}
	// aes解密
	decryptData, err := decrypt(encryptDataByte, []byte(key))
	if err != nil {
		return "", err
	}
	return string(decryptData), nil
}

func EncryptUserAddress(rawData, key string) (string, error) {
	// aes加密
	encryptData, err := encrypt([]byte(rawData), []byte(key))
	if err != nil {
		return "", err
	}
	// base64编码
	return base64.StdEncoding.EncodeToString(encryptData), nil
}

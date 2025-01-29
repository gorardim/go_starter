package avater

import (
	"app/pkg/randx"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (c *Client) GetAvaterUrl(email string) (string, error) {
	md5email := md5.Sum([]byte(email))
	md5emailStr := fmt.Sprintf("%x", md5email)
	httpResponse, err := c.HttpClient.Get(c.Url + md5emailStr + "?s=200&d=identicon&r=PG")
	if err != nil {
		return "", err
	}
	defer httpResponse.Body.Close()
	now := time.Now()
	filename := fmt.Sprintf("/images/%s/%s.png", now.Format("20060102"), randx.Seq(16))
	dst := c.Storage + filename
	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return "", err
	}
	file, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, httpResponse.Body)
	if err != nil {
		return "", err
	}
	return c.CdnUrl + filename, nil
}

func (c *Client) SaveBase64Tofile(base64Img string) (string, error) {
	now := time.Now()
	filename := fmt.Sprintf("/images/%s/%s.png", now.Format("20060102"), randx.Seq(16))
	dst := c.Storage + filename
	err := os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return "", err
	}
	file, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer file.Close()
	src := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Img))
	_, err = io.Copy(file, src)
	if err != nil {
		return "", err
	}
	return "assets/" + filename, nil
}

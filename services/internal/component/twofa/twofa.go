package twofa

import (
	"app/services/internal/component/apiuser"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"net/url"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	Issuer        string = "inTrip"
	Period        string = "30"
	AlgorithmSHA1 string = "SHA1"
	Digits        string = "6"
)

type TwoFA struct{}

type TwoFAResponse struct {
	QrCode string `json:"qr_code"`
	Secret string `json:"secret"`
}

func getBase64QrCode(qrImage image.Image) string {
	var qrPng bytes.Buffer
	png.Encode(&qrPng, qrImage)
	return base64.StdEncoding.EncodeToString(qrPng.Bytes())
}

func (t *TwoFA) Generate(ctx context.Context) (*TwoFAResponse, error) {
	user := apiuser.MustFromContext(ctx)
	var secret *otp.Key
	var err error
	if user.SecretTwoFA == "" {
		if secret, err = totp.Generate(totp.GenerateOpts{
			Issuer:      Issuer,
			AccountName: user.Username,
		}); err != nil {
			return nil, err
		}
	} else {
		v := url.Values{}
		v.Set("algorithm", AlgorithmSHA1)
		v.Set("digits", Digits)
		v.Set("issuer", Issuer)
		v.Set("period", Period)
		v.Set("secret", user.SecretTwoFA)

		u := url.URL{
			Scheme:   "otpauth",
			Host:     "totp",
			Path:     "/" + Issuer + ":" + user.Username + "?",
			RawQuery: v.Encode(),
		}

		fmt.Println("check Url patern ===", u.String())

		if secret, err = otp.NewKeyFromURL(u.String()); err != nil {
			return nil, err
		}
	}
	qrImage, err := secret.Image(200, 200)
	if err != nil {
		return nil, err
	}
	return &TwoFAResponse{
		Secret: secret.Secret(),
		QrCode: getBase64QrCode(qrImage),
	}, nil
}

func (t *TwoFA) Verify(ctx context.Context, passcode string, secret string) bool {
	return totp.Validate(passcode, secret)
}

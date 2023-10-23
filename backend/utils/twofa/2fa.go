package twofa

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type TwoFA struct {
	Issuer      string
	Email       string
	Key         *otp.Key
	ImageBase64 string
}

func (t *TwoFA) NewKey() error {
	var err error
	t.Key, err = totp.Generate(totp.GenerateOpts{
		Issuer:      t.Issuer,
		AccountName: t.Email,
	})
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	img, err := t.Key.Image(200, 200)
	png.Encode(&buf, img)
	t.ImageBase64 = base64.StdEncoding.EncodeToString(buf.Bytes())

	if err != nil {
		return err
	}
	return nil
}

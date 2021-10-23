package utils

import (
	"bytes"
	"encoding/base64"
	"github.com/pquerna/otp/totp"
	"go-auth/models/orm"
	_struct "go-auth/models/struct"
	"image/png"
)

func GenerateB64Qr(data orm.User) _struct.QrData {
	if data.TwoFactEnabled {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      data.Name,
			AccountName: data.Email,
		})

		if err != nil {
			panic(err)
		}
		var buf bytes.Buffer
		img, err := key.Image(200, 200)
		if err != nil {
			panic(err)
		}
		encodingErr := png.Encode(&buf, img)

		if encodingErr != nil {
			panic(encodingErr)
		}
		return _struct.QrData{
			TwoFactEnabled: true,
			Secret:         key.Secret(),
			QrCode:         base64.StdEncoding.EncodeToString(buf.Bytes()),
		}
	}
	return _struct.QrData{
		TwoFactEnabled: false,
		Secret:         "",
		QrCode:         "",
	}
}

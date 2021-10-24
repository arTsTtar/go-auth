package utils

import (
	"bytes"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
	"go-auth/models/dto/request"
	_struct "go-auth/models/struct"
	"image/png"
	"strconv"
	"time"
)

func GenerateB64Qr(data request.UserRequest) _struct.QrData {
	if data.TwoFactEnabled {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      *data.Name,
			AccountName: *data.Email,
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

func CreateAuthCookie(userId uint, secret string) (fiber.Cookie, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userId)),
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
	})

	token, err := claims.SignedString([]byte(secret))

	if err != nil {
		return fiber.Cookie{}, err
	}

	return fiber.Cookie{
		Name:     "jwtToken",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 5),
		HTTPOnly: true,
	}, nil

}
func CreateBackupCodes() [6]string {
	var backupCodes [6]string
	backupCodes[0] = GenerateRandomPasswd()
	backupCodes[1] = GenerateRandomPasswd()
	backupCodes[2] = GenerateRandomPasswd()
	backupCodes[3] = GenerateRandomPasswd()
	backupCodes[4] = GenerateRandomPasswd()
	backupCodes[5] = GenerateRandomPasswd()
	return backupCodes
}

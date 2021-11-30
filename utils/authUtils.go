package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
	"go-auth/entity"
	"go-auth/models/dto/request"
	_struct "go-auth/models/struct"
	"image/png"
	"strconv"
	"time"
)

const SecretKey = "adsfadsfasdfnuasnfuias23as98fasj8dfjas/asdfiijasdf"

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

func CreateAuthCookieAndHandleError(user *entity.User, minutes time.Duration) (*fiber.Cookie, error, int) {
	cookie, err := CreateAuthCookie(user, SecretKey, minutes)

	if err != nil {
		return nil, err, 500
	}

	return &cookie, err, 200

}

func CreateAuthCookie(user *entity.User, secret string, minutes time.Duration) (fiber.Cookie, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(user.Id)),
		"Roles":     user.Roles,
		"ExpiresAt": time.Now().Add(time.Minute * minutes).Unix(),
	})

	token, err := claims.SignedString([]byte(secret))

	if err != nil {
		return fiber.Cookie{}, err
	}

	return fiber.Cookie{
		Name:     "jwtToken",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * minutes),
		HTTPOnly: true,
	}, nil

}

func CheckAuthenticationFromCookie(c *fiber.Ctx) (*jwt.Token, error) {
	cookie := c.Cookies("jwtToken")
	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, errors.New("unauthenticated")
	}
	return token, nil
}
func CreateBackupCodes() [6]string {
	var backupCodes [6]string
	backupCodes[0] = *GenerateRandomPasswd()
	backupCodes[1] = *GenerateRandomPasswd()
	backupCodes[2] = *GenerateRandomPasswd()
	backupCodes[3] = *GenerateRandomPasswd()
	backupCodes[4] = *GenerateRandomPasswd()
	backupCodes[5] = *GenerateRandomPasswd()
	return backupCodes
}

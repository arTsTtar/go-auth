package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
	"go-auth/db"
	"go-auth/models"
	"go-auth/response"
	"golang.org/x/crypto/bcrypt"
	"image/png"
	"strconv"
	"time"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      data["name"],
		AccountName: data["email"],
	})
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	secret, _ := bcrypt.GenerateFromPassword([]byte(key.Secret()), 14)

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

	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
		Secret:   secret,
	}
	db.DB.Create(&user)

	userResponse := response.UserResponse{
		Id:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Secret: key.Secret(),
		QrCode: imgBase64Str,
	}
	return c.JSON(userResponse)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	db.DB.Where("email = ?", data["email"]).First(&user)

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))

	if err != nil || user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Wrong credentials",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not log in.",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwtToken",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 5),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Successfully logged in!",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwtToken")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	db.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwtToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Minute * 5),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Successfully logged out.",
	})
}

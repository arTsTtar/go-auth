package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-auth/db"
	"go-auth/models/dto/request"
	"go-auth/models/dto/response"
	"go-auth/models/entity"
	"go-auth/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const SecretKey = "adsfadsfasdfnuasnfuias23as98fasj8dfjas/asdfiijasdf"

func Register(c *fiber.Ctx) error {
	var data request.UserRequest

	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return err
	}

	var existingUser entity.User
	db.DB.Where("email = ?", data.Email).First(&existingUser)

	if existingUser.Id != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User with this email already exists",
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 12)

	var backupPasswords = utils.CreateBackupCodes()

	var qrData = utils.GenerateB64Qr(data)

	user := entity.User{
		Name:           data.Name,
		Email:          data.Email,
		Password:       password,
		TwoFactEnabled: qrData.TwoFactEnabled,
		TwoFactSecret:  qrData.Secret,
	}
	db.DB.Create(&user)

	for i := 0; i < len(backupPasswords); i++ {
		backupPasswd, _ := bcrypt.GenerateFromPassword([]byte(backupPasswords[i]), 12)
		backupCode := entity.BackupCode{
			UserId:     user.Id,
			BackupCode: backupPasswd,
		}
		db.DB.Create(&backupCode)
	}

	userResponse := response.UserCreationResponse{
		Id:             user.Id,
		Name:           user.Name,
		Email:          user.Email,
		TwoFactEnabled: qrData.TwoFactEnabled,
		Secret:         qrData.Secret,
		QrCode:         qrData.QrCode,
		BackupCodes:    backupPasswords,
	}
	return c.JSON(userResponse)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user entity.User

	db.DB.Where("email = ?", data["email"]).First(&user)

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))

	if err != nil || user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Wrong credentials",
		})
	}

	cookie, err := utils.CreateAuthCookie(user.Id, SecretKey)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not log in.",
		})
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

	var user entity.User

	db.DB.Where("id = ?", claims.Issuer).First(&user)

	userResponse := response.SimpleUserResponse{
		Id:             user.Id,
		Name:           user.Name,
		Email:          user.Email,
		TwoFactEnabled: user.TwoFactEnabled,
	}

	return c.JSON(userResponse)
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

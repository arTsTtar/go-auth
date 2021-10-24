package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-auth/db"
	"go-auth/models/dto/response"
	"go-auth/models/entity"
)

type UserService interface {
	User(c *fiber.Ctx)
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

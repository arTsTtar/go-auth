package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-auth/services"
)

type UserController interface {
	User(c *fiber.Ctx) error
}

type userController struct {
	authService services.AuthService
}

func NewUerController(as services.AuthService) UserController {
	return userController{
		authService: as,
	}
}

func (u userController) User(c *fiber.Ctx) error {
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

	userResponse, err := u.authService.GetUserDetailsFromToken(token)
	return c.JSON(userResponse)
}

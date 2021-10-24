package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-auth/models/dto/request"
	"go-auth/services"
)

type AuthController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(s services.AuthService) AuthController {
	return authController{
		authService: s,
	}
}

const SecretKey = "adsfadsfasdfnuasnfuias23as98fasj8dfjas/asdfiijasdf"

func (a authController) Register(c *fiber.Ctx) error {
	var data request.UserRequest
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return err
	}

	if !request.ValidateCreation(data) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Bad request.",
		})
	}
	userResponse, err := a.authService.Register(data)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(userResponse)
}

func (a authController) Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	cookie, err, status := a.authService.Login(data)

	if err != nil {
		c.Status(status)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Cookie(cookie)

	c.Status(status)
	return c.JSON(fiber.Map{
		"message": "Successfully logged in!",
	})
}

func (a authController) Logout(c *fiber.Ctx) error {
	cookie := a.authService.Logout()
	c.Cookie(cookie)
	return c.JSON(fiber.Map{
		"message": "Successfully logged out.",
	})
}

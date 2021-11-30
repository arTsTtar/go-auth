package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-auth/models/dto/request"
	"go-auth/models/dto/response"
	"go-auth/services"
	"go-auth/utils"
)

type UserController interface {
	User(c *fiber.Ctx) error
	ChangePasswordAndUpdate2FA(c *fiber.Ctx) error
}

type userController struct {
	authService services.AuthService
}

func NewUserController(as services.AuthService) UserController {
	return userController{
		authService: as,
	}
}

func (u userController) User(c *fiber.Ctx) error {

	token, err := utils.CheckAuthenticationFromCookie(c)

	if err != nil && err.Error() == "unauthenticated" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims := token.Claims.(jwt.MapClaims)

	userResponse, err := u.authService.GetUserDetails(claims["Issuer"].(string))
	return c.JSON(userResponse)
}

func (u userController) ChangePasswordAndUpdate2FA(c *fiber.Ctx) error {
	token, err := utils.CheckAuthenticationFromCookie(c)

	if err != nil && err.Error() == "unauthenticated" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var data request.ChangePassword
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return err
	}

	validationError := request.ValidateChangePassword(data)
	if validationError != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": validationError.Error(),
		})
	}

	claims := token.Claims.(jwt.MapClaims)

	user, err := u.authService.ChangePassword(claims["Issuer"].(string), data)

	if data.Disable2FA {
		err = u.authService.Disable2FA(user)
	}
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(response.ToSimpleUserResponse(user))
}

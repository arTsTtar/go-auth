package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-auth/models/dto/response"
	"go-auth/services"
	"go-auth/utils"
)

type AdminController interface {
	GetUserList(c *fiber.Ctx) error
	ResetUserData(c *fiber.Ctx) error
}

type adminController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAdminController(as services.AuthService, us services.UserService) AdminController {
	return adminController{
		authService: as,
		userService: us,
	}
}

func (a adminController) GetUserList(c *fiber.Ctx) error {
	token, err := utils.CheckAuthenticationFromCookie(c)

	if err != nil && err.Error() == "unauthenticated" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	var issuer = claims["Issuer"].(string)

	isAdmin, err := a.authService.CheckIfUserIsAdmin(issuer)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if !isAdmin {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized access for the resource",
		})
	}

	users, err := a.userService.GetAllUsers()
	var usersResponse response.SimpleUserResponseArray
	for _, user := range users {
		usersResponse = append(usersResponse, response.ToSimpleUserResponse(user))
	}
	return c.JSON(usersResponse)
}

func (a adminController) ResetUserData(c *fiber.Ctx) error {
	panic("implement me")
}

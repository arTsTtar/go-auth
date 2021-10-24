package services

import (
	"github.com/gofiber/fiber/v2"
	"go-auth/repository"
	"go-auth/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type authService struct {
	backupCodeRepository repository.BackupCodeRepository
	userRepository       repository.UserRepository
}

func NewAuthService(ur repository.UserRepository, bcr repository.BackupCodeRepository) AuthService {
	return authService{
		userRepository:       ur,
		backupCodeRepository: bcr,
	}
}

const SecretKey = "adsfadsfasdfnuasnfuias23as98fasj8dfjas/asdfiijasdf"

func (a authService) Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := a.userRepository.FindUserByEmail(data["email"])

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))

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

func (a authService) Register(c *fiber.Ctx) error {
	panic("implement me")
}

func (a authService) Logout(c *fiber.Ctx) error {
	panic("implement me")
}

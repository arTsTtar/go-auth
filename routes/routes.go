package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-auth/controllers"
	"go-auth/db"
	"go-auth/repository"
	"go-auth/services"
)

func Setup(app *fiber.App) {
	userRepository := repository.NewUserRepository(db.DB)
	backupCodeRepository := repository.NewBackupCodeRepository(db.DB)
	authService := services.NewAuthService(userRepository, backupCodeRepository)
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUerController(authService)

	app.Post("/api/register", authController.Register)
	app.Post("/api/login", authController.Login)
	app.Post("/api/altLogin", authController.AltLogin)
	app.Get("/api/user", userController.User)
	app.Post("/api/logout", authController.Logout)
}

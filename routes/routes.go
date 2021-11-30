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
	roleRepository := repository.NewRoleRepository(db.DB)
	backupCodeRepository := repository.NewBackupCodeRepository(db.DB)
	authService := services.NewAuthService(userRepository, roleRepository, backupCodeRepository)
	userService := services.NewUserService(userRepository, authService)
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(authService)
	adminController := controllers.NewAdminController(authService, userService)

	// ADMIN ENDPOINTS:
	app.Get("/api/admin/users", adminController.GetUserList)

	// Auth Endpoints (User)
	app.Post("/api/register", authController.Register)
	app.Post("/api/login", authController.Login)
	app.Post("/api/altLogin", authController.AltLogin)
	app.Post("/api/logout", authController.Logout)

	//user endpoints
	app.Get("/api/user", userController.User)
	app.Post("/api/user/changePassword", userController.ChangePasswordAndUpdate2FA)
	app.Post("/api/user/:id/resetPassword", userController.ResetToRandomPassword)
}

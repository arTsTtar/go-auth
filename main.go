package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-auth/db"
	"go-auth/routes"
)

func main() {
	db.Connect()
	db.Automigrate()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	port := ":8000"
	routes.Setup(app)
	err := app.Listen(port)
	if err != nil {
		panic("could not listen to port " + port)
		return
	}
}

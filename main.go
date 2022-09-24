// @/main.go
package main

import (
	"fliteklub/config"
	"fliteklub/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func addRoutes(app *fiber.App) {
	userHandler := handlers.UserHandler{}
	userHandler.AddRoutes(app)
}

func main() {
	app := fiber.New()

	err := config.Connect()
	if err != nil {
		return
	}

	addRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

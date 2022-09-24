// @/main.go
package main

import (
	"fliteklub/config"
	"fliteklub/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func sayHello(c *fiber.Ctx) error {
	return c.Status(200).JSON("hello world!")
}

func addRoutes(app *fiber.App) {
	app.Get("/", sayHello)
	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.AddUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.RemoveUser)
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

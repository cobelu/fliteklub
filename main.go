// @/main.go
package main

import (
	"fliteklub/config"
	"fliteklub/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	err := config.Connect()
	if err != nil {
		return
	}

	addRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func addRoutes(app *fiber.App) {
	routeHandlers := []handlers.CrudHandler{
		handlers.UserHandler{},
		handlers.AircraftHandler{},
		handlers.ClubHandler{},
	}

	for _, h := range routeHandlers {
		h.AddRoutes(app)
		h.AddExamples()
	}
}

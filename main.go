// @/main.go
package main

import (
	"fliteklub/config"
	"fliteklub/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	var err error

	app := fiber.New()

	err = config.Connect()
	if err != nil {
		return
	}

	err = config.Cleanup()
	if err != nil {
		return
	}

	err = config.Migrate()
	if err != nil {
		return
	}

	addRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func addRoutes(app *fiber.App) {
	routeHandlers := []handlers.CrudHandler{
		handlers.AircraftHandler{},
		handlers.ClubHandler{},
		handlers.MembershipHandler{},
		handlers.OwnershipHandler{},
		handlers.ReservationHandler{},
		handlers.UserHandler{},
	}

	for _, h := range routeHandlers {
		h.AddRoutes(app)
		h.AddExamples()
	}
}

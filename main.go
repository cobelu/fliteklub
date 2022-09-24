// @/main.go
package main

import (
	"fliteklub/config"
	"fliteklub/handlers"
	"fliteklub/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	var err error

	app := fiber.New()
	models := models()

	err = config.Connect()
	if err != nil {
		return
	}

	err = config.Cleanup(models)
	if err != nil {
		return
	}

	err = config.Migrate(models)
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
		handlers.ReservationHandler{},
		handlers.UserHandler{},
	}

	for _, h := range routeHandlers {
		h.AddRoutes(app)
		h.AddExamples()
	}
}

func models() []model.Model {
	aircraft := model.Aircraft{}
	club := model.Club{}
	reservation := model.Reservation{}
	user := model.User{}

	return []model.Model{
		aircraft,
		club,
		reservation,
		user,
	}
}

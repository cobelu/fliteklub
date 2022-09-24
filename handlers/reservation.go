package handlers

import (
	"fliteklub/config"
	"fliteklub/model"
	"github.com/gofiber/fiber/v2"
)

type ReservationHandler struct {
	c Crudder
}

func (h ReservationHandler) getAll(c *fiber.Ctx) error {
	var reservations []model.Reservation

	config.Database.Find(&reservations)

	return c.Status(200).JSON(reservations)
}

func (h ReservationHandler) get(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation model.Reservation

	result := config.Database.Find(&reservation, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(reservation)
}

func (h ReservationHandler) insert(c *fiber.Ctx) error {
	reservation := new(model.Reservation)

	if err := c.BodyParser(reservation); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&reservation)

	return c.Status(201).JSON(reservation)
}

func (h ReservationHandler) update(c *fiber.Ctx) error {
	reservation := new(model.Reservation)
	id := c.Params("id")

	if err := c.BodyParser(reservation); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&reservation)

	return c.Status(200).JSON(reservation)
}

func (h ReservationHandler) delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation model.Reservation

	result := config.Database.Delete(&reservation, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func (h ReservationHandler) AddRoutes(app *fiber.App) {
	app.Get("/reservations", h.getAll)
	app.Get("/reservations/:id", h.get)
	app.Post("/reservations", h.insert)
	app.Put("/reservations/:id", h.update)
	app.Delete("/reservations/:id", h.delete)
}

func (h ReservationHandler) AddExamples() {
	example := model.ExampleReservation()
	config.Database.Create(&example)
}

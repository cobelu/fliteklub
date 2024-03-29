package handlers

import (
	"fliteklub/config"
	"fliteklub/model"
	"github.com/gofiber/fiber/v2"
)

type AircraftHandler struct {
	c Crudder
}

func (h AircraftHandler) getAll(c *fiber.Ctx) error {
	var aircrafts []model.Aircraft

	config.Database.Find(&aircrafts)

	return c.Status(200).JSON(aircrafts)
}

func (h AircraftHandler) get(c *fiber.Ctx) error {
	id := c.Params("id")
	var aircraft model.Aircraft

	result := config.Database.Find(&aircraft, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(aircraft)
}

func (h AircraftHandler) insert(c *fiber.Ctx) error {
	aircraft := new(model.Aircraft)

	if err := c.BodyParser(aircraft); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&aircraft)

	return c.Status(201).JSON(aircraft)
}

func (h AircraftHandler) update(c *fiber.Ctx) error {
	aircraft := new(model.Aircraft)
	id := c.Params("id")

	if err := c.BodyParser(aircraft); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&aircraft)

	return c.Status(200).JSON(aircraft)
}

func (h AircraftHandler) delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var aircraft model.Aircraft

	result := config.Database.Delete(&aircraft, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func (h AircraftHandler) AddRoutes(app *fiber.App) {
	app.Get("/aircrafts", h.getAll)
	app.Get("/aircrafts/:id", h.get)
	app.Post("/aircrafts", h.insert)
	app.Put("/aircrafts/:id", h.update)
	app.Delete("/aircrafts/:id", h.delete)
}

func (h AircraftHandler) AddExamples() {
	example := model.ExampleAircraft()
	config.Database.Create(&example)
}

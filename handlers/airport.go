package handlers

import (
	"fliteklub/config"
	"fliteklub/model"
	"github.com/gofiber/fiber/v2"
)

type AirportHandler struct {
	c Crudder
}

func (h AirportHandler) getAll(c *fiber.Ctx) error {
	var airports []model.Airport

	config.Database.Find(&airports)

	return c.Status(200).JSON(airports)
}

func (h AirportHandler) get(c *fiber.Ctx) error {
	id := c.Params("id")
	var airport model.Airport

	dbModel := config.Database.Model(&model.Airport{})
	result := dbModel.Find(&airport, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(airport)
}

func (h AirportHandler) getByName(c *fiber.Ctx) error {
	name := c.Params("name")
	var airport model.Airport

	dbModel := config.Database.Model(&model.Airport{})
	// TODO How to do conditions?
	result := dbModel.Find(&airport, "name = ?", name)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(airport)
}

func (h AirportHandler) AddRoutes(app *fiber.App) {
	app.Get("/airports", h.getAll)
	app.Get("/airports/:id", h.get)
}

func (h AirportHandler) AddExamples() {
	example := model.ExampleAirport()
	config.Database.Create(&example)
}

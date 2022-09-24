package handlers

import (
	"fliteklub/config"
	"fliteklub/model"
	"github.com/gofiber/fiber/v2"
)

type OwnershipHandler struct {
	c Crudder
}

func (h OwnershipHandler) get(c *fiber.Ctx) error {
	var ownerships []model.Ownership

	config.Database.Find(&ownerships)

	return c.Status(200).JSON(ownerships)
}

func (h OwnershipHandler) getAll(c *fiber.Ctx) error {
	id := c.Params("id")
	var ownership model.Ownership

	result := config.Database.Find(&ownership, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(ownership)
}

func (h OwnershipHandler) insert(c *fiber.Ctx) error {
	ownership := new(model.Ownership)

	if err := c.BodyParser(ownership); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&ownership)

	return c.Status(201).JSON(ownership)
}

func (h OwnershipHandler) update(c *fiber.Ctx) error {
	ownership := new(model.Ownership)
	id := c.Params("id")

	if err := c.BodyParser(ownership); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&ownership)

	return c.Status(200).JSON(ownership)
}

func (h OwnershipHandler) delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var ownership model.Ownership

	result := config.Database.Delete(&ownership, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func (h OwnershipHandler) AddRoutes(app *fiber.App) {
	app.Get(h.c.Path("/ownership/:id"), h.getAll)
	app.Get("/ownership/:id", h.get)
	app.Post("/ownership", h.insert)
	app.Put("/ownership/:id", h.update)
	app.Delete("/ownership/:id", h.delete)
}

func (h OwnershipHandler) AddExamples() {
	example := model.ExampleOwnership()
	config.Database.Create(&example)
}

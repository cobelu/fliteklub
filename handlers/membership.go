package handlers

import (
	"fliteklub/config"
	"fliteklub/model"
	"github.com/gofiber/fiber/v2"
)

type MembershipHandler struct {
	c Crudder
}

func (h MembershipHandler) get(c *fiber.Ctx) error {
	var memberships []model.Membership

	config.Database.Find(&memberships)

	return c.Status(200).JSON(memberships)
}

func (h MembershipHandler) getAll(c *fiber.Ctx) error {
	id := c.Params("id")
	var membership model.Membership

	result := config.Database.Find(&membership, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(membership)
}

func (h MembershipHandler) insert(c *fiber.Ctx) error {
	membership := new(model.Membership)

	if err := c.BodyParser(membership); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&membership)

	return c.Status(201).JSON(membership)
}

func (h MembershipHandler) update(c *fiber.Ctx) error {
	membership := new(model.Membership)
	id := c.Params("id")

	if err := c.BodyParser(membership); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&membership)

	return c.Status(200).JSON(membership)
}

func (h MembershipHandler) delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var membership model.Membership

	result := config.Database.Delete(&membership, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func (h MembershipHandler) AddRoutes(app *fiber.App) {
	app.Get(h.c.Path("/membership/:id"), h.getAll)
	app.Get("/membership/:id", h.get)
	app.Post("/membership", h.insert)
	app.Put("/membership/:id", h.update)
	app.Delete("/membership/:id", h.delete)
}

func (h MembershipHandler) AddExamples() {
	example := model.ExampleOwnership()
	config.Database.Create(&example)
}

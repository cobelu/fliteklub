package handlers

import (
	"fliteklub/config"
	"fliteklub/model"
	"github.com/gofiber/fiber/v2"
)

type ClubHandler struct {
	c Crudder
}

func (h ClubHandler) getAll(c *fiber.Ctx) error {
	var clubs []model.Club

	config.Database.Find(&clubs)

	return c.Status(200).JSON(clubs)
}

func (h ClubHandler) get(c *fiber.Ctx) error {
	id := c.Params("id")
	var club model.Club

	dbModel := config.Database.Model(&model.Club{})
	preloaded := dbModel.Preload("memberships").Preload("ownerships")
	result := preloaded.Find(&club, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(club)
}

func (h ClubHandler) insert(c *fiber.Ctx) error {
	club := new(model.Club)

	if err := c.BodyParser(club); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&club)

	return c.Status(201).JSON(club)
}

func (h ClubHandler) update(c *fiber.Ctx) error {
	club := new(model.Club)
	id := c.Params("id")

	if err := c.BodyParser(club); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&club)

	return c.Status(200).JSON(club)
}

func (h ClubHandler) delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var club model.Club

	result := config.Database.Delete(&club, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func (h ClubHandler) AddRoutes(app *fiber.App) {
	app.Get("/clubs", h.getAll)
	app.Get("/clubs/:id", h.get)
	app.Post("/clubs", h.insert)
	app.Put("/clubs/:id", h.update)
	app.Delete("/clubs/:id", h.delete)
}

func (h ClubHandler) AddExamples() {
	example := model.ExampleClub()
	config.Database.Create(&example)
}

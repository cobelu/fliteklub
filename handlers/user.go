package handlers

import (
	"fliteklub/config"
	"fliteklub/model"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	c Crudder
}

func (h UserHandler) getAll(c *fiber.Ctx) error {
	var users []model.User

	config.Database.Find(&users)

	return c.Status(200).JSON(users)
}

func (h UserHandler) get(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User

	result := config.Database.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(user)
}

func (h UserHandler) insert(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&user)

	return c.Status(201).JSON(user)
}

func (h UserHandler) update(c *fiber.Ctx) error {
	user := new(model.User)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&user)

	return c.Status(200).JSON(user)
}

func (h UserHandler) delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User

	result := config.Database.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func (h UserHandler) AddRoutes(app *fiber.App) {
	app.Get("/users", h.getAll)
	app.Get("/users/:id", h.get)
	app.Post("/users", h.insert)
	app.Put("/users/:id", h.update)
	app.Delete("/users/:id", h.delete)
}

func (h UserHandler) AddExamples() {
	example := model.ExampleUser
	config.Database.Create(&example)
}

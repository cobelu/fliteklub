package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type CrudHandler interface {
	insert(c *fiber.Ctx) error
	get(c *fiber.Ctx) error
	getAll(c *fiber.Ctx) error
	update(c *fiber.Ctx) error
	delete(c *fiber.Ctx) error
	AddRoutes(app *fiber.App)
}

type Crudder struct {
	name string
}

func (c Crudder) Path(url string) string {
	return fmt.Sprintf(url, c.name)
}

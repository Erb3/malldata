package presenters

import "github.com/gofiber/fiber/v2"

func PlotCreated(id string) *fiber.Map {
	return &fiber.Map{
		"id": id,
	}
}

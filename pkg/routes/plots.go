package routes

import (
	"log"

	"github.com/Pixium/MallData/pkg/db"
	"github.com/Pixium/MallData/pkg/presenters"
	"github.com/Pixium/MallData/pkg/types"
	"github.com/gofiber/fiber/v2"
)

func NewPlot(c *fiber.Ctx) error {
	data := new(types.Plot)

	if err := c.BodyParser(data); err != nil {
		log.Fatal("Error while creating new plot: ", err)
		return c.SendStatus(500)
	}

	id, err := db.InsertPlot(data)
	if err != nil {
		log.Fatal("Error while creating new plot: ", err)
		return c.SendStatus(500)
	}

	return c.JSON(presenters.PlotCreated(id))
}

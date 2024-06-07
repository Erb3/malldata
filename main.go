package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/Pixium/MallData/pkg/db"
	"github.com/Pixium/MallData/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed all:frontend/dist/*
var frontend_fs embed.FS

func main() {
	_, err := db.SetupDatabase()
	if err != nil {
		log.Fatalf("unable to setup database due to %s", err.Error())
	}

	app := fiber.New()
	api := app.Group("/api")

	frontend_dir, err := fs.Sub(frontend_fs, "frontend/dist")
	if err != nil {
		log.Panicf("unable to sub frontend build directory: %s. Did you forget to compile the frontend?", err.Error())
	}

	app.Use(filesystem.New(filesystem.Config{
		Root: http.FS(frontend_dir),
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} ${latency} - ${method} ${path}\n",
	}))

	api.Post("/plot", routes.NewPlot)

	log.Fatal(app.Listen(":3019"))
}

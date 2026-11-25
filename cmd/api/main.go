package main

import (
	"log"

	"github.com/bn-lima/goBlog/internal/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

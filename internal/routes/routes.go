package routes

import (
	"github.com/bn-lima/goBlog/internal/handler"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("main page")
	})

	app.Post("/user/register", handler.RegisterUser)
}

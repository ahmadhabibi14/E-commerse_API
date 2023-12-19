package app

import (
	"e-commerse_api/handler"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	app.Get("/product/:id", handler.GetProduct)
}

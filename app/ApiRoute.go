package app

import (
	"e-commerse_api/handler"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	app.Get("/products/:id", handler.GetProduct)
	app.Post("/products", handler.CreateProduct)
	app.Delete("/products/:id", handler.DeleteProduct)
	app.Get("/products", handler.GetProductLists)
}

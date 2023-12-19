package handler

import (
	"context"
	"database/sql"
	"e-commerse_api/conf"
	"e-commerse_api/models/data"
	"e-commerse_api/models/web"

	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectDB()
	ctx := context.Background()
	defer db.Close()

	id := c.Params(`id`)
	product := data.NewProduct(db)
	productRow, err := product.FindById(ctx, id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: STATUS_NOTFOUND,
			Data:   productRow,
		}
		return c.Status(fiber.StatusNotFound).JSON(webResponse)
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: STATUS_OK,
		Data:   productRow,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

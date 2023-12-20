package handler

import (
	"context"
	"database/sql"
	"e-commerse_api/conf"
	"e-commerse_api/models/data"
	"e-commerse_api/models/web"
	"e-commerse_api/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func DeleteProduct(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectDB()
	ctx := context.Background()
	defer db.Close()

	id := c.Params(`id`)
	msg, err := utils.ValidateStruct(web.ProductGetRequest{Id: id})
	if err != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_INVALIDPAYLOAD,
			Data:   msg,
		}
		return c.Status(fiber.StatusBadRequest).JSON(webResponse)
	}

	product := data.NewProduct(db)
	_, productErr := product.FindById(ctx, fmt.Sprintf("%s", id))
	if productErr != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_NOTFOUND,
			Data:   `Product not found`,
		}
		return c.Status(fiber.StatusBadRequest).JSON(webResponse)
	}

	product.Delete(ctx, id)

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: STATUS_OK,
		Data:   `Product deleted`,
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusCreated).JSON(webResponse)
}

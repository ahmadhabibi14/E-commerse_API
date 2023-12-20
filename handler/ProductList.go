package handler

import (
	"context"
	"database/sql"
	"e-commerse_api/conf"
	"e-commerse_api/models/data"
	"e-commerse_api/models/web"

	"github.com/gofiber/fiber/v2"
)

func GetProductLists(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectDB()
	ctx := context.Background()
	defer db.Close()

	product := data.NewProduct(db)
	productList, err := product.FindAll(ctx)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: STATUS_SERVERERROR,
			Data:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(webResponse)
	}

	if productList == nil {
		productList = []web.ProductListResponse{}
	}
	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: STATUS_OK,
		Data:   productList,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

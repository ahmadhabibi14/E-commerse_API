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

	sort := c.Query("sortBy")
	product := data.NewProduct(db)

	if sort != `` {
		if sort == `title` {
			productList, err := product.FindAllSortByTitle(ctx)
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
		} else if sort == `rating` {
			productList, err := product.FindAllSortByRating(ctx)
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
	}

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
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

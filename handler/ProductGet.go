package handler

import (
	"database/sql"
	"e-commerse_api/conf"
	"e-commerse_api/models/data"
	"e-commerse_api/models/web"
	"e-commerse_api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {
	zlog := conf.InitLogger()
	var db *sql.DB = conf.ConnectDB()
	ctx := c.Context()
	defer db.Close()

	id := c.Params(`id`)
	msg, err := utils.ValidateStruct(web.ProductGetRequest{Id: id})
	if err != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_INVALIDPAYLOAD,
			Errors: msg[0],
			Data:   ``,
		}
		zlog.Error().Str("ERROR", err.Error()).Msg("Error validate product ID")
		return c.Status(fiber.StatusBadRequest).JSON(webResponse)
	}
	product := data.NewProduct(db)
	productRow, err := product.FindById(ctx, id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: STATUS_NOTFOUND,
			Errors: `Product not found`,
			Data:   ``,
		}
		zlog.Log().Str("ERROR", err.Error()).Msg("Product not found")
		return c.Status(fiber.StatusNotFound).JSON(webResponse)
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: STATUS_OK,
		Errors: ``,
		Data:   productRow,
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

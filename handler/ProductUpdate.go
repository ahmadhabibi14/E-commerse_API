package handler

import (
	"database/sql"
	"e-commerse_api/conf"
	"e-commerse_api/models/data"
	"e-commerse_api/models/web"
	"e-commerse_api/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UpdateProduct(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectDB()
	ctx := c.Context()
	defer db.Close()

	id := c.Params(`id`)
	var in web.ProductUpdateRequest
	if err := c.BodyParser(&in); err != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: `Invalid input`,
			Data:   ``,
		}
		return c.Status(fiber.StatusBadRequest).JSON(webResponse)
	}
	in.Id = id

	msg, err := utils.ValidateStruct(in)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_INVALIDPAYLOAD,
			Errors: msg[0],
			Data:   ``,
		}
		return c.Status(fiber.StatusBadRequest).JSON(webResponse)
	}

	product := data.NewProduct(db)
	_, productErr := product.FindById(ctx, fmt.Sprintf("%s", id))
	if productErr != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_NOTFOUND,
			Errors: `Product not found`,
			Data:   ``,
		}
		return c.Status(fiber.StatusBadRequest).JSON(webResponse)
	}

	productUpdateReq := web.ProductUpdateRequest{
		Id:          in.Id,
		Title:       in.Title,
		Description: in.Description,
		Rating:      in.Rating,
		Image:       in.Image,
	}
	productUpdateErr := product.Update(ctx, productUpdateReq)
	if productUpdateErr != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_INVALIDPAYLOAD,
			Errors: productUpdateErr.Error(),
			Data:   ``,
		}
		return c.Status(fiber.StatusBadRequest).JSON(webResponse)
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: STATUS_OK,
		Errors: ``,
		Data:   `Product updated`,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

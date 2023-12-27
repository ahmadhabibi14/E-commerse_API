package handler

import (
	"database/sql"
	"e-commerse_api/conf"
	"e-commerse_api/models/data"
	"e-commerse_api/models/web"
	"e-commerse_api/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateProduct(c *fiber.Ctx) error {
	var db *sql.DB = conf.ConnectDB()
	ctx := c.Context()
	defer db.Close()

	var in web.ProductNewRequest
	if err := c.BodyParser(&in); err != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: `Invalid input`,
			Data:   ``,
		}
		return c.Status(fiber.StatusBadRequest).JSON(webResponse)
	}

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

	id := uuid.New()
	productData := web.ProductNewRequest{
		Id:          fmt.Sprintf("%v", id),
		Title:       in.Title,
		Description: in.Description,
		Rating:      in.Rating,
		Image:       in.Image,
	}
	product := data.NewProduct(db)
	if productErr := product.Insert(ctx, productData); productErr != nil {
		webResponse := web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: `Cannot create product`,
			Data:   ``,
		}
		return c.Status(fiber.StatusBadRequest).JSON(webResponse)
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: STATUS_OK,
		Errors: ``,
		Data:   `Product created`,
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusCreated).JSON(webResponse)
}

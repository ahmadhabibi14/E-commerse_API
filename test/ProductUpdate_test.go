package test

import (
	"bytes"
	"e-commerse_api/conf"
	"e-commerse_api/handler"
	"e-commerse_api/models/web"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	zlog := conf.InitLogger()
	err := godotenv.Load("../.env")
	if err != nil {
		zlog.Error().
			Str("ERROR", err.Error()).
			Msg("cannot load .env files")
	}
}

func TestProductUpdate(t *testing.T) {
	app := fiber.New()
	app.Put(`/products`, handler.UpdateProduct)

	t.Run("Product_Updated", func(t *testing.T) {
		id := `1ff96e76-bc8b-42be-9ce9-d000aace5d54`
		payload := web.ProductUpdateRequest{
			Id:          id,
			Title:       `Title to be update`,
			Description: `Description to be update`,
			Rating:      4.8,
			Image:       `/path/to/image.png`,
		}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			t.Error(`Failed to convert JSON data`)
		}

		req := httptest.NewRequest(http.MethodPut, `/products`, bytes.NewBuffer(jsonData))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode, `Product updated`)

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})

	t.Run("Product_Invalid", func(t *testing.T) {
		payload := web.ProductNewRequest{
			Id:          `invalid-id`,
			Title:       `random stuff`,
			Description: `desc`,
			Rating:      5.0,
			Image:       `idk-what-is-this`,
		}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			t.Error(`Failed to convert JSON data`)
		}

		req := httptest.NewRequest(http.MethodPost, `/products`, bytes.NewBuffer(jsonData))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, `Invalid request`)

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})
}

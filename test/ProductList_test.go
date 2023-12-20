package test

import (
	"e-commerse_api/conf"
	"e-commerse_api/handler"
	"io"
	"net/http/httptest"
	"testing"

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

func TestProductList(t *testing.T) {
	app := fiber.New()
	app.Get(`/products`, handler.GetProductLists)

	t.Run("Products_Found", func(t *testing.T) {
		req := httptest.NewRequest(fiber.MethodGet, `/products`, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode, `Products found`)

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})

	t.Run("Products_Sort_by_Title", func(t *testing.T) {
		req := httptest.NewRequest(fiber.MethodGet, `/products?sortBy=title`, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode, `Products found`)

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})

	t.Run("Products_Sort_by_Rating", func(t *testing.T) {
		req := httptest.NewRequest(fiber.MethodGet, `/products?sortBy=rating`, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode, `Products found`)

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})
}

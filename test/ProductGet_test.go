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

func TestProductGet(t *testing.T) {
	app := fiber.New()
	app.Get(`/products/:id`, handler.GetProduct)

	t.Run("Product Found", func(t *testing.T) {
		id := `1ff96e76-bc8b-42be-9ce9-d000aace5d54`
		req := httptest.NewRequest(fiber.MethodGet, `/products/`+id, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode, `Product found`)

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})

	t.Run("Product Not Found", func(t *testing.T) {
		id := `this-is-invalid-id-xxxxxxxxxxxxxxxx`
		req := httptest.NewRequest(fiber.MethodGet, `/products/`+id, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode, `Product Not found`)

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})

	t.Run("Invalid Payload/Request", func(t *testing.T) {
		id := `this-cannot-be-validated`
		req := httptest.NewRequest(fiber.MethodGet, `/products/`+id, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, `Product ID invalid`)

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})
}

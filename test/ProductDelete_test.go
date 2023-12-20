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

func TestProductDelete(t *testing.T) {
	app := fiber.New()
	app.Delete(`/products/:id`, handler.DeleteProduct)

	t.Run("Product_Deleted", func(t *testing.T) {
		id := `1ff96e76-bc8b-42be-9ce9-d000aace5d54`
		req := httptest.NewRequest(fiber.MethodDelete, `/products/`+id, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusCreated, resp.StatusCode, `Product deleted`)

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})

	t.Run("Product_Delete_not_found", func(t *testing.T) {
		id := `this-is-invalid-id-xxxxxxxxxxxxxxxxx`
		req := httptest.NewRequest(fiber.MethodDelete, `/products/`+id, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, `Product not found`)

		body, _ := io.ReadAll(resp.Body)
		t.Logf(`Response: %v`, string(body))
	})
}

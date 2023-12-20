package test

import (
	"e-commerse_api/handler"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestProductList(t *testing.T) {
	app := fiber.New()
	app.Get(`/products`, handler.GetProductLists)

	req := httptest.NewRequest(fiber.MethodGet, `/products`, nil)
	resp, _ := app.Test(req)

	// assert.Equal(t, fiber.StatusNotFound, resp.StatusCode, `Not found`)

	body, _ := io.ReadAll(resp.Body)
	t.Logf(`Response: %v`, string(body))
}

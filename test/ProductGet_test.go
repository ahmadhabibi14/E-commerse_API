package test

import (
	"e-commerse_api/handler"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestProductGet(t *testing.T) {
	app := fiber.New()
	app.Get(`/product/:id`, handler.GetProduct)

	id := `this-must-be-the-actual-id`

	req := httptest.NewRequest(fiber.MethodGet, `/product/`+id, nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode, `Not found`)

	body, _ := io.ReadAll(resp.Body)
	t.Logf(`Response: %v`, string(body))
}

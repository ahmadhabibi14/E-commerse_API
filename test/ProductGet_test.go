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
	app.Get(`/products/:id`, handler.GetProduct)

	id := `1ff96e76-bc8b-42be-9ce9-d000aace5d54`

	req := httptest.NewRequest(fiber.MethodGet, `/products/`+id, nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode, `Not found`)

	body, _ := io.ReadAll(resp.Body)
	t.Logf(`Response: %v`, string(body))
}

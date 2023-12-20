package test

import (
	"bytes"
	"e-commerse_api/conf"
	"e-commerse_api/handler"
	"e-commerse_api/models/web"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func TestProductNew(t *testing.T) {
	app := fiber.New()
	app.Post(`/products`, handler.CreateProduct)

	t.Run("Product_Created", func(t *testing.T) {
		id := uuid.New()
		payload := web.ProductNewRequest{
			Id:    fmt.Sprintf("%v", id),
			Title: `Lenovo Thinkpad x280`,
			Description: `Laptop Lenovo Thinkpad X280
			Kondisi Oke, normal, Siap pakai. Mesin ori fisik 90 - 98% masih bagus dan Mulus
			
			Spesifikasi:
			Intel Core i7 -8350U @1.7 Ghz (8 Cpus)~1.9ghz
			RAM DDR4 16 gb
			SSD 1 TB
			Wifi
			Camera/webcam
			Hdmi port
			USB port
			VGA intel
			Layar 12.5inch jernih no death pixel
			Layar FHD
			Layar Touchscreen jari
			Windows 10/11 mozilla/game/office/anti virus.
			
			kondisi mesin ok/prima
			engsel kokoh
			fisik super mulus
			speaker jernih dan normal
			baterai ori tahan 1-3jam
			adaptor
			garansi 1 bulan full tuker unit langsung`,
			Rating: 5.0,
			Image:  `/assets/img/lenovo-thinkpad-x280.png`,
		}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			t.Error(`Failed to convert JSON data`)
		}

		req := httptest.NewRequest(http.MethodPost, `/products`, bytes.NewBuffer(jsonData))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusCreated, resp.StatusCode, `Product created`)

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

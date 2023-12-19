package app

import (
	"e-commerse_api/conf"
	"e-commerse_api/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type WebServer struct {
	AppName string
	Cfg     conf.WebConf
}

func (w *WebServer) Start() {
	app := fiber.New(fiber.Config{
		AppName: w.AppName,
		Prefork: false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				`error`: `ini error bang`,
			})
		},
	})
	app.Use(requestid.New())
	app.Use(logger.New(middlewares.LoggerConfig))
	app.Use(limiter.New(middlewares.Limiter))
	app.Use(cors.New(middlewares.CORSConfig))

	ApiRoutes(app)

	log.Fatal(app.Listen(w.Cfg.ListenAddr()))
}

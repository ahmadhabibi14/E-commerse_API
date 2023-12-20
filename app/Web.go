package app

import (
	"e-commerse_api/conf"
	"e-commerse_api/handler"
	"e-commerse_api/middlewares"
	"e-commerse_api/models/web"
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
			webResponse := web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: handler.STATUS_NOTFOUND,
				Data:   `Resource not found`,
			}
			return c.Status(fiber.StatusNotFound).JSON(webResponse)
		},
	})
	app.Use(requestid.New())
	app.Use(logger.New(middlewares.LoggerConfig))
	app.Use(limiter.New(middlewares.Limiter))
	app.Use(cors.New(middlewares.CORSConfig))

	ApiRoutes(app)

	log.Fatal(app.Listen(w.Cfg.ListenAddr()))
}

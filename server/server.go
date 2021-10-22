package server

import (
	"github.com/felipeolivers/gofiber-projectbase-api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/redirect/v2"
)

func Start() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       config.Env.AppName})

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/swagger/index.html",
		},
		StatusCode: 301,
	}))
	app.Get("/dashboard", monitor.New())

	loadMiddleware(app)
	loadRoutes(app)

	app.Listen(":" + config.Env.AppPortServer)

}

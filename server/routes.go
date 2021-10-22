package server

import (
	"github.com/felipeolivers/gofiber-projectbase-api/server/routes"
	"github.com/gofiber/fiber/v2"
)

func loadRoutes(app *fiber.App) {
	routes.Healthz(app)
	routes.Swagger(app)
}

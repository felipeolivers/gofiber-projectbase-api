package routes

import (
	"github.com/felipeolivers/gofiber-projectbase-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Healthz(app *fiber.App) {
	g := app.Group("/healthz")
	g.Get("", controllers.Healthz)
}

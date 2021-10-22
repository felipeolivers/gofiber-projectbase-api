package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/felipeolivers/gofiber-projectbase-api/config"
	"github.com/felipeolivers/gofiber-projectbase-api/docs"
	"github.com/gofiber/fiber/v2"
)

// @description Copyright © 2021, Luiz Filipe Miranda de Oliveira.
// @description Todos os direitos reservados.
func Swagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.Handler)

	// Swagger Project Base API
	docs.SwaggerInfo.Title = config.Env.AppName
	docs.SwaggerInfo.Version = config.Env.AppVersion
}

package controllers

import (
	"github.com/felipeolivers/gofiber-projectbase-api/config"
	"github.com/felipeolivers/gofiber-projectbase-api/models"
	"github.com/felipeolivers/gofiber-projectbase-api/tools"
	"github.com/gofiber/fiber/v2"
)

// @Tags healthz
// @Summary Get Api Info
// @Description Get Api Info
// @Accept json
// @Produce json
// @Success 200 {object} models.Healthz
// @Router /healthz [get]
func Healthz(c *fiber.Ctx) error {
	span, _ := tools.ApmSpan(c.Context(), tools.Self())
	defer span.End()

	info := models.Healthz{
		Name:        config.Env.AppName,
		Version:     config.Env.AppVersion,
		Environment: config.Env.AppEnvironment,
	}
	return c.JSON(info)
}

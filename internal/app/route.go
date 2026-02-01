package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/addixit1/fiber-boilerplate/internal/modules/user/v1"
)

func registerRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	userv1.Routes(api)
}

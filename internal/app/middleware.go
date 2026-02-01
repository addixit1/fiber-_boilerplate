package app

import (
	"github.com/addixit1/fiber-boilerplate/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func registerMiddlewares(app *fiber.App) {

	// PANIC RECOVERY
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true, // dev me useful
	}))

	// REQUEST HEADERS (Platform, Timezone, Language, etc.)
	app.Use(middleware.RequestHeaders())

	// DETAILED REQUEST LOGGER
	app.Use(middleware.DetailedLogger())

}

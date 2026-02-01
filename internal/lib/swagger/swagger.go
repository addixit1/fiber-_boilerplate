package swagger

import (
	_ "github.com/addixit1/fiber-boilerplate/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Register(app *fiber.App) {
	app.Get("/swagger/*", swagger.New(swagger.Config{
		Title: "User Services",
	}))
}

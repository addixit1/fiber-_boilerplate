package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/addixit1/fiber-boilerplate/internal/config"

)

func BasicAuth() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			config.BASIC_USERNAME: config.BASIC_PASSWORD,
		},
		Realm: "Restricted",
	})
}
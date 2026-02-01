package userv1

import (
	"github.com/addixit1/fiber-boilerplate/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router) {
	r.Post("/users", middleware.BasicAuth(), Create)
	r.Get("/users", middleware.BasicAuth(), List)
}

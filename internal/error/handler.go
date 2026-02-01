package errors

import (
	"github.com/addixit1/fiber-boilerplate/internal/config"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, err error) error {

	// Default status
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// PROD MODE (safe)
	if config.Config.Env == "production" {
		return c.Status(code).JSON(fiber.Map{
			"message": eMessage(err),
		})
	}

	// DEV MODE (safe debug)
	response := fiber.Map{
		"error":  err.Error(),
		"path":   c.Path(),
		"method": c.Method(),
		"query":  c.Queries(),
		"body":   string(c.Body()),
	}

	// SAFE PARAM ACCESS
	if c.Route() != nil {
		response["params"] = c.AllParams()
	}

	return c.Status(code).JSON(response)
}

func eMessage(err error) string {
	if e, ok := err.(*fiber.Error); ok {
		return e.Message
	}
	return "Something went wrong"
}

package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Language middleware extracts language from request headers
func Language() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Try to get language from Accept-Language header
		lang := c.Get("Accept-Language", "en")

		// If empty or not supported, default to English
		if lang == "" {
			lang = "en"
		}

		// Take only first 2 characters (e.g., "en-US" -> "en")
		if len(lang) > 2 {
			lang = lang[:2]
		}

		// Store language in context for use in handlers
		c.Locals("lang", lang)

		return c.Next()
	}
}

package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// RequestHeaders middleware extracts common headers from request
func RequestHeaders() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract and store headers in context

		// Language (already handled by Language middleware, but keeping for completeness)
		lang := c.Get("accept-language", "en")
		if len(lang) > 2 {
			lang = lang[:2]
		}
		c.Locals("lang", lang)

		// Platform: 1-Android, 2-iOS, 3-WEB
		platform := c.Get("platform", "3")
		c.Locals("platform", platform)

		// Timezone
		timezone := c.Get("timezone", "Asia/Kolkata")
		c.Locals("timezone", timezone)

		// Timezone Offset
		offset := c.Get("offset", "0")
		c.Locals("offset", offset)

		// App Version
		appversion := c.Get("appversion", "v1")
		c.Locals("appversion", appversion)

		// Route Version
		routeversion := c.Get("routeversion", "v1")
		c.Locals("routeversion", routeversion)

		return c.Next()
	}
}

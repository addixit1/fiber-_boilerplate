package middleware

import (
	"github.com/addixit1/fiber-boilerplate/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)


func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return fiber.ErrUnauthorized
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(config.Config.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			return fiber.ErrUnauthorized
		}

		return c.Next()
	}
}

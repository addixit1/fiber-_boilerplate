package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// DetailedLogger prints detailed request information to terminal
func DetailedLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Capture start time
		start := time.Now()

		// Process request first
		err := c.Next()

		// Only log if not an unauthorized attempt (skip 401 responses)
		// This prevents duplicate logs for Basic Auth retry
		if c.Response().StatusCode() == 401 {
			return err
		}

		// Print request start
		fmt.Printf("\n\n%s\n", "----------REQUEST STARTS----------")

		// Print request type and path
		fmt.Printf("::1: %s %s\n", c.Method(), c.Path())
		fmt.Printf("Request Type======> %s\n", c.Method())
		fmt.Printf("Request Path======> %s\n", c.Path())

		// Print request body
		body := c.Body()
		if len(body) > 0 {
			var bodyMap map[string]interface{}
			if err := json.Unmarshal(body, &bodyMap); err == nil {
				bodyJSON, _ := json.MarshalIndent(bodyMap, "  ", "  ")
				fmt.Printf("Request Body======> %s\n", string(bodyJSON))
			} else {
				fmt.Printf("Request Body======> %s\n", string(body))
			}
		} else {
			fmt.Printf("Request Body======> {}\n")
		}

		// Print request params
		fmt.Printf("Request Params=====> {}\n")

		// Print request query
		queryMap := make(map[string]string)
		c.Request().URI().QueryArgs().VisitAll(func(key, value []byte) {
			queryMap[string(key)] = string(value)
		})
		if len(queryMap) > 0 {
			queryJSON, _ := json.MarshalIndent(queryMap, "  ", "  ")
			fmt.Printf("Request Query=====> %s\n", string(queryJSON))
		} else {
			fmt.Printf("Request Query=====> {Object: null prototype} {}\n")
		}

		// Print authorization header
		auth := c.Get("Authorization", "")
		if auth != "" {
			if len(auth) > 6 {
				fmt.Printf("Authorization=====> %s\n", auth)
			}
		}

		// Print API key
		apiKey := c.Get("api_key", "1234")
		fmt.Printf("api_key===========> %s\n", apiKey)

		// Print headers
		fmt.Printf("platform==========> %s\n", c.Get("platform", "1"))
		fmt.Printf("timezone==========> %s\n", c.Get("timezone", "undefined"))
		fmt.Printf("offset============> %s\n", c.Get("offset", "undefined"))
		fmt.Printf("language==========> %s\n", c.Get("accept-language", "en"))

		// Print response status
		fmt.Printf("\nResponse Status===> %d\n", c.Response().StatusCode())

		// Print request end
		duration := time.Since(start)
		fmt.Printf("%s\n", "----------REQUEST ENDS----------")
		fmt.Printf("Completed in: %v\n\n", duration)

		return err
	}
}

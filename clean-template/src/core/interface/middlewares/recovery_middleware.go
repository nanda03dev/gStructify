package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Custom recovery middleware for Fiber
func RecoveryMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Defer the recovery logic to catch any panics
		defer func() {
			if err := recover(); err != nil {
				// Log the panic error for further investigation
				log.Printf("Recovered from panic: %v", err)

				// Return a generic Internal Server Error response
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Internal Server Error",
				})
			}
		}()

		// Continue with the next handler
		return c.Next()
	}
}

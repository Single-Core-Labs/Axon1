// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Bypass health check
		if c.Path() == "/health" {
			return c.Next()
		}

		apiKey := c.Get("Authorization")
		if apiKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		// In a real system, we'd validate against Postgres/Redis here
		return c.Next()
	}
}

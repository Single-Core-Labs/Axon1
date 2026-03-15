// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	if err := InitDB(); err != nil {
		log.Fatalf("DB Init failed: %v", err)
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Auth routes
	auth := app.Group("/auth")
	auth.Post("/signup", SignupHandler)
	auth.Post("/login", LoginHandler)
	auth.Get("/me", JWTMiddleware(), MeHandler)

	// API routes
	api := app.Group("/v1", JWTMiddleware())
	
	api.Get("/projects", ListProjectsHandler)
	api.Post("/projects", CreateProjectHandler)
	api.Get("/projects/:id/keys", ListAPIKeysHandler)
	api.Post("/projects/:id/keys", CreateAPIKeyHandler)
	api.Delete("/keys/:keyId", RevokeAPIKeyHandler)

	api.Get("/traces", ListTracesHandler)
	api.Get("/stats", StatsHandler)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8081"
	}

	log.Fatal(app.Listen(":" + port))
}

// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg, err := LoadConfig(context.Background(), "axon.yaml")
	if err != nil {
		log.Printf("Failed to load config, using defaults: %v", err)
		cfg = &Config{}
		cfg.Server.Port = 8080
	}

	tracer, _ := NewTracer(context.Background(), os.Getenv("CLICKHOUSE_URL"))
	
	router := NewRouter(cfg, tracer)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api := app.Group("/v1", AuthMiddleware())

	api.Post("/chat/completions", func(c *fiber.Ctx) error {
		var req ChatRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
		}

		ctx, cancel := context.WithTimeout(c.Context(), 30*time.Second)
		defer cancel()

		resp, err := router.Route(ctx, req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(resp)
	})

	// Setup graceful shutdown
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gateway...")
	app.Shutdown()
}

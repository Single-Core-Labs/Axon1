// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TestHealthCheck(t *testing.T) {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	req := httptest.NewRequest("GET", "/health", nil)
	resp, _ := app.Test(req, -1)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestRouterFallback(t *testing.T) {
	cfg := &Config{}
	cfg.Strategies.Fallback = []StrategyConfig{
		{Provider: "openai", Model: "gpt-4"},
	}
	
	r := NewRouter(cfg, nil)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	resp, err := r.Route(ctx, ChatRequest{Prompt: "hello"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !strings.Contains(resp.Completion, "hello") {
		t.Errorf("Expected completion to contain hello, got %s", resp.Completion)
	}
}

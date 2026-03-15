// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"context"
	"net/http"
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
	// Mock OpenAI server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"choices":[{"message":{"content":"hello response"}}]}`))
	}))
	defer server.Close()

	cfg := &Config{}
	cfg.Strategies.Fallback = []StrategyConfig{
		{Provider: "openai", Model: "gpt-4"},
	}
	
	r := NewRouter(cfg, nil)
	// Override the OpenAI provider's baseURL to point to our mock server
	if p, ok := r.providers["openai"].(*OpenAIProvider); ok {
		p.baseURL = server.URL
	}

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

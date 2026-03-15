// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"context"
	"net/http"
	"time"
)

type Provider interface {
	Complete(ctx context.Context, req ChatRequest, model string) (*ChatResponse, error)
	Name() string
}

type BaseProvider struct {
	name   string
	apiKey string
	client *http.Client
}

func (p *BaseProvider) Name() string {
	return p.name
}

// Simple mock implementation for OpenAI adapter
type OpenAIProvider struct {
	BaseProvider
}

func NewOpenAIProvider(apiKey string) *OpenAIProvider {
	return &OpenAIProvider{
		BaseProvider: BaseProvider{
			name:   "openai",
			apiKey: apiKey,
			client: &http.Client{Timeout: 30 * time.Second},
		},
	}
}

func (p *OpenAIProvider) Complete(ctx context.Context, req ChatRequest, model string) (*ChatResponse, error) {
	// Mock returning a response
	start := time.Now()
	// simulate latency
	time.Sleep(10 * time.Millisecond)

	return &ChatResponse{
		Provider:   p.name,
		Model:      model,
		Completion: "Response to: " + req.Prompt,
		LatencyMS:  uint32(time.Since(start).Milliseconds()),
	}, nil
}

// Similar mock adapters for others
type AnthropicProvider struct { BaseProvider }
type GeminiProvider struct { BaseProvider }
type MistralProvider struct { BaseProvider }

func NewAnthropicProvider(apiKey string) *AnthropicProvider {
	return &AnthropicProvider{BaseProvider{name: "anthropic", apiKey: apiKey, client: &http.Client{Timeout: 30 * time.Second}}}
}
func (p *AnthropicProvider) Complete(ctx context.Context, req ChatRequest, model string) (*ChatResponse, error) {
	start := time.Now()
	return &ChatResponse{Provider: p.name, Model: model, Completion: "Claude says: " + req.Prompt, LatencyMS: uint32(time.Since(start).Milliseconds())}, nil
}

func NewGeminiProvider(apiKey string) *GeminiProvider {
	return &GeminiProvider{BaseProvider{name: "gemini", apiKey: apiKey, client: &http.Client{Timeout: 30 * time.Second}}}
}
func (p *GeminiProvider) Complete(ctx context.Context, req ChatRequest, model string) (*ChatResponse, error) {
	start := time.Now()
	return &ChatResponse{Provider: p.name, Model: model, Completion: "Gemini says: " + req.Prompt, LatencyMS: uint32(time.Since(start).Milliseconds())}, nil
}

func NewMistralProvider(apiKey string) *MistralProvider {
	return &MistralProvider{BaseProvider{name: "mistral", apiKey: apiKey, client: &http.Client{Timeout: 30 * time.Second}}}
}
func (p *MistralProvider) Complete(ctx context.Context, req ChatRequest, model string) (*ChatResponse, error) {
	start := time.Now()
	return &ChatResponse{Provider: p.name, Model: model, Completion: "Mistral says: " + req.Prompt, LatencyMS: uint32(time.Since(start).Milliseconds())}, nil
}

// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Provider interface {
	Complete(ctx context.Context, req ChatRequest, model string) (*ChatResponse, error)
	Name() string
}

type BaseProvider struct {
	name    string
	apiKey  string
	baseURL string
	client  *http.Client
}

func (p *BaseProvider) Name() string {
	return p.name
}

// OpenAI provider structures
type openaiChatRequest struct {
	Model    string          `json:"model"`
	Messages []openaiMessage `json:"messages"`
}

type openaiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openaiChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error"`
}

// Real implementation for OpenAI adapter
type OpenAIProvider struct {
	BaseProvider
}

func NewOpenAIProvider(apiKey string) *OpenAIProvider {
	return &OpenAIProvider{
		BaseProvider: BaseProvider{
			name:    "openai",
			apiKey:  apiKey,
			baseURL: "https://api.openai.com/v1",
			client:  &http.Client{Timeout: 60 * time.Second},
		},
	}
}

func (p *OpenAIProvider) Complete(ctx context.Context, req ChatRequest, model string) (*ChatResponse, error) {
	start := time.Now()

	oaReq := openaiChatRequest{
		Model: model,
		Messages: []openaiMessage{
			{Role: "user", Content: req.Prompt},
		},
	}

	body, err := json.Marshal(oaReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.baseURL+"/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("openai request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var oaErr openaiChatResponse
		json.Unmarshal(respBody, &oaErr)
		if oaErr.Error != nil {
			return nil, fmt.Errorf("openai error (%d): %s", resp.StatusCode, oaErr.Error.Message)
		}
		return nil, fmt.Errorf("openai error status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	var oaResp openaiChatResponse
	if err := json.Unmarshal(respBody, &oaResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(oaResp.Choices) == 0 {
		return nil, fmt.Errorf("no choices returned from openai")
	}

	return &ChatResponse{
		Provider:   p.name,
		Model:      model,
		Completion: oaResp.Choices[0].Message.Content,
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

// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenAIProvider_Complete(t *testing.T) {
	// Create a mock server to act as OpenAI API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify path
		if r.URL.Path != "/chat/completions" {
			t.Errorf("Expected /chat/completions path, got %s", r.URL.Path)
		}

		// Verify headers
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("Expected Bearer test-key auth header, got %s", r.Header.Get("Authorization"))
		}

		// Send mock response
		resp := openaiChatResponse{
			Choices: []struct {
				Message struct {
					Content string `json:"content"`
				} `json:"message"`
			}{
				{
					Message: struct {
						Content string `json:"content"`
					}{
						Content: "This is a mock completion",
					},
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	p := NewOpenAIProvider("test-key")
	p.baseURL = server.URL // Inject mock server URL

	resp, err := p.Complete(context.Background(), ChatRequest{Prompt: "hello"}, "text-davinci-003")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.Completion != "This is a mock completion" {
		t.Errorf("Expected mock completion, got %s", resp.Completion)
	}
}

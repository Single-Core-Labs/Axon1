// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"time"
)

type Trace struct {
	TraceID          string
	ProjectID        string
	Model            string
	Provider         string
	Prompt           string
	Completion       string
	PromptTokens     uint32
	CompletionTokens uint32
	TotalTokens      uint32
	CostUSD          float64
	LatencyMS        uint32
	Status           string
	ErrorMsg         string
	CreatedAt        time.Time
}

type ChatRequest struct {
	Prompt string `json:"prompt"`
}

type ChatResponse struct {
	Provider   string `json:"provider"`
	Model      string `json:"model"`
	Completion string `json:"completion"`
	LatencyMS  uint32 `json:"latency_ms"`
}

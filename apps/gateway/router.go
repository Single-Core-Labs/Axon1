// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Router struct {
	providers map[string]Provider
	config    *Config
	tracer    *Tracer
}

func NewRouter(cfg *Config, tracer *Tracer) *Router {
	r := &Router{
		providers: make(map[string]Provider),
		config:    cfg,
		tracer:    tracer,
	}

	// Initialize providers (normally with env vars)
	r.providers["openai"] = NewOpenAIProvider("sk-mock")
	r.providers["anthropic"] = NewAnthropicProvider("sk-mock")
	r.providers["gemini"] = NewGeminiProvider("sk-mock")
	r.providers["mistral"] = NewMistralProvider("sk-mock")

	return r
}

func (r *Router) Route(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	var lastErr error

	// Fallback strategy 
	for _, strat := range r.config.Strategies.Fallback {
		p, ok := r.providers[strat.Provider]
		if !ok {
			continue
		}

		start := time.Now()
		resp, err := p.Complete(ctx, req, strat.Model)
		latency := uint32(time.Since(start).Milliseconds())

		trace := Trace{
			TraceID:      uuid.NewString(),
			ProjectID:    "proj-123", // mocked project id
			Model:        strat.Model,
			Provider:     strat.Provider,
			Prompt:       req.Prompt,
			CostUSD:      0.001,
			LatencyMS:    latency,
			Status:       "success",
			CreatedAt:    time.Now(),
		}

		if err != nil {
			lastErr = err
			trace.Status = "error"
			trace.ErrorMsg = err.Error()
			
			if r.tracer != nil {
				r.tracer.Record(trace)
			}
			continue // try next fallback
		}

		trace.Completion = resp.Completion
		if r.tracer != nil {
			r.tracer.Record(trace)
		}

		return resp, nil
	}

	if lastErr != nil {
		return nil, fmt.Errorf("all fallback providers failed: %w", lastErr)
	}
	return nil, fmt.Errorf("no providers configured")
}

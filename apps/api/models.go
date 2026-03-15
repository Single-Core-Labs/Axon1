// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	OrgID        uuid.UUID `json:"org_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Project struct {
	ID        uuid.UUID `json:"id"`
	OrgID     uuid.UUID `json:"org_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type APIKey struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
	Name      string    `json:"name"`
	KeyHash   string    `json:"-"`
	Preview   string    `json:"preview"`
	Revoked   bool      `json:"revoked"`
	CreatedAt time.Time `json:"created_at"`
}

type Trace struct {
	TraceID          string    `json:"trace_id"`
	ProjectID        string    `json:"project_id"`
	Model            string    `json:"model"`
	Provider         string    `json:"provider"`
	Prompt           string    `json:"prompt"`
	Completion       string    `json:"completion"`
	PromptTokens     uint32    `json:"prompt_tokens"`
	CompletionTokens uint32    `json:"completion_tokens"`
	TotalTokens      uint32    `json:"total_tokens"`
	CostUSD          float64   `json:"cost_usd"`
	LatencyMS        uint32    `json:"latency_ms"`
	Status           string    `json:"status"`
	ErrorMsg         string    `json:"error_msg"`
	CreatedAt        time.Time `json:"created_at"`
}

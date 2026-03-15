// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type Tracer struct {
	db     clickhouse.Conn
	traces chan Trace
}

func NewTracer(ctx context.Context, url string) (*Tracer, error) {
	opts, err := clickhouse.ParseDSN(url)
	if err != nil {
		return nil, fmt.Errorf("invalid clickhouse url: %w", err)
	}

	conn, err := clickhouse.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to clickhouse: %w", err)
	}

	if err := conn.Ping(ctx); err != nil {
		// Log warning but don't fail, might be starting up in docker-compose
		fmt.Printf("clickhouse not ready yet: %v\n", err)
	}

	t := &Tracer{
		db:     conn,
		traces: make(chan Trace, 10000), // Buffered channel to never block
	}

	go t.worker()
	return t, nil
}

func (t *Tracer) Record(trace Trace) {
	select {
	case t.traces <- trace:
	default:
		// Queue is full, drop trace to avoid blocking request (in prod to avoid memory leak)
		fmt.Println("trace dropped, queue full")
	}
}

func (t *Tracer) worker() {
	var batch []Trace
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case trace, ok := <-t.traces:
			if !ok {
				t.flush(batch)
				return
			}
			batch = append(batch, trace)
			if len(batch) >= 1000 {
				t.flush(batch)
				batch = nil
			}
		case <-ticker.C:
			if len(batch) > 0 {
				t.flush(batch)
				batch = nil
			}
		}
	}
}

func (t *Tracer) flush(batch []Trace) {
	if len(batch) == 0 {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	b, err := t.db.PrepareBatch(ctx, "INSERT INTO traces")
	if err != nil {
		return // log error in real system
	}

	for _, trace := range batch {
		err := b.Append(
			trace.TraceID,
			trace.ProjectID,
			trace.Model,
			trace.Provider,
			trace.Prompt,
			trace.Completion,
			trace.PromptTokens,
			trace.CompletionTokens,
			trace.TotalTokens,
			trace.CostUSD,
			trace.LatencyMS,
			trace.Status,
			trace.ErrorMsg,
			trace.CreatedAt,
		)
		if err != nil {
			continue
		}
	}
	_ = b.Send() // ignore error for robustness
}

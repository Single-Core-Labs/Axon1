// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ListTracesHandler(c *fiber.Ctx) error {
	projectID := c.Query("project_id")
	if projectID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "project_id is required"})
	}

	query := "SELECT trace_id, project_id, model, provider, prompt, completion, total_tokens, cost_usd, latency_ms, status, created_at FROM traces WHERE project_id = ? ORDER BY created_at DESC LIMIT 100"
	
	rows, err := CH.Query(context.Background(), query, projectID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("failed to fetch traces: %v", err)})
	}
	defer rows.Close()

	var traces []Trace
	for rows.Next() {
		var t Trace
		err := rows.Scan(&t.TraceID, &t.ProjectID, &t.Model, &t.Provider, &t.Prompt, &t.Completion, &t.TotalTokens, &t.CostUSD, &t.LatencyMS, &t.Status, &t.CreatedAt)
		if err != nil {
			continue
		}
		traces = append(traces, t)
	}

	return c.JSON(traces)
}

func StatsHandler(c *fiber.Ctx) error {
	projectID := c.Query("project_id")
	if projectID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "project_id is required"})
	}

	// Simple aggregate stats for now
	query := "SELECT count(), sum(cost_usd), avg(latency_ms) FROM traces WHERE project_id = ?"
	var count uint64
	var cost float64
	var avgLatency float64

	err := CH.QueryRow(context.Background(), query, projectID).Scan(&count, &cost, &avgLatency)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to calculate stats"})
	}

	return c.JSON(fiber.Map{
		"total_traces": count,
		"total_cost":   cost,
		"avg_latency":  avgLatency,
	})
}

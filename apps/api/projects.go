// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/xyproto/randomstring"
)

func ListProjectsHandler(c *fiber.Ctx) error {
	orgID := c.Locals("org_id").(string)
	rows, err := DB.Query("SELECT id, org_id, name, created_at FROM projects WHERE org_id = $1", orgID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch projects"})
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var p Project
		rows.Scan(&p.ID, &p.OrgID, &p.Name, &p.CreatedAt)
		projects = append(projects, p)
	}
	return c.JSON(projects)
}

func CreateProjectHandler(c *fiber.Ctx) error {
	orgID := c.Locals("org_id").(string)
	var req struct{ Name string `json:"name"` }
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	id := uuid.New()
	_, err := DB.Exec("INSERT INTO projects (id, org_id, name) VALUES ($1, $2, $3)", id, orgID, req.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create project"})
	}

	return c.Status(201).JSON(fiber.Map{"id": id, "name": req.Name})
}

func CreateAPIKeyHandler(c *fiber.Ctx) error {
	projectID := c.Params("id")
	var req struct{ Name string `json:"name"` }
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	key := "ax_" + randomstring.HumanFriendlyString(32)
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(key)))
	preview := key[:7] + "..." + key[len(key)-4:]

	id := uuid.New()
	_, err := DB.Exec("INSERT INTO api_keys (id, project_id, name, key_hash, preview) VALUES ($1, $2, $3, $4, $5)",
		id, projectID, req.Name, hash, preview)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create API key"})
	}

	// Return plaintext once
	return c.Status(201).JSON(fiber.Map{
		"id":      id,
		"name":    req.Name,
		"key":     key,
		"preview": preview,
	})
}

func ListAPIKeysHandler(c *fiber.Ctx) error {
	projectID := c.Params("id")
	rows, err := DB.Query("SELECT id, project_id, name, preview, revoked, created_at FROM api_keys WHERE project_id = $1", projectID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch API keys"})
	}
	defer rows.Close()

	var keys []APIKey
	for rows.Next() {
		var k APIKey
		rows.Scan(&k.ID, &k.ProjectID, &k.Name, &k.Preview, &k.Revoked, &k.CreatedAt)
		keys = append(keys, k)
	}
	return c.JSON(keys)
}

func RevokeAPIKeyHandler(c *fiber.Ctx) error {
	id := c.Params("keyId")
	_, err := DB.Exec("UPDATE api_keys SET revoked = true WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to revoke API key"})
	}
	return c.SendStatus(204)
}

// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	OrgName  string `json:"org_name"`
}

func SignupHandler(c *fiber.Ctx) error {
	var req SignupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	orgID := uuid.New()
	userID := uuid.New()

	// In a real app, this would be a transaction
	_, err := DB.Exec("INSERT INTO organizations (id, name) VALUES ($1, $2)", orgID, req.OrgName)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create organization"})
	}

	_, err = DB.Exec("INSERT INTO users (id, email, password_hash, org_id) VALUES ($1, $2, $3, $4)",
		userID, req.Email, string(hash), orgID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.JSON(fiber.Map{"message": "User created successfully"})
}

func LoginHandler(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user User
	err := DB.QueryRow("SELECT id, email, password_hash, org_id FROM users WHERE email = $1", req.Email).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.OrgID)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"org_id":  user.OrgID.String(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "supersecret"
	}

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": t})
}

func MeHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var user User
	err := DB.QueryRow("SELECT id, email, org_id FROM users WHERE id = $1", userID).
		Scan(&user.ID, &user.Email, &user.OrgID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

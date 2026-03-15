// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/lib/pq"
)

var (
	DB   *sql.DB
	CH   clickhouse.Conn
)

func InitDB() error {
	var err error
	pgUrl := os.Getenv("DATABASE_URL")
	if pgUrl == "" {
		pgUrl = "postgres://axon:axon@localhost:5432/axon?sslmode=disable"
	}

	DB, err = sql.Open("postgres", pgUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err := DB.Ping(); err != nil {
		fmt.Printf("Warning: Postgres not ready: %v\n", err)
	}

	chUrl := os.Getenv("CLICKHOUSE_URL")
	if chUrl == "" {
		chUrl = "http://localhost:8123"
	}

	opts, err := clickhouse.ParseDSN(chUrl)
	if err != nil {
		return fmt.Errorf("failed to parse clickhouse dsn: %w", err)
	}

	CH, err = clickhouse.Open(opts)
	if err != nil {
		return fmt.Errorf("failed to connect to clickhouse: %w", err)
	}

	if err := CH.Ping(context.Background()); err != nil {
		fmt.Printf("Warning: ClickHouse not ready: %v\n", err)
	}

	return nil
}

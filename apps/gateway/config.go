// Copyright (c) 2026 1Core Labs. MIT License.
package main

import (
	"context"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ProviderConfig struct {
	BaseURL string `yaml:"base_url"`
}

type StrategyConfig struct {
	Provider string `yaml:"provider"`
	Model    string `yaml:"model"`
}

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Providers  map[string]ProviderConfig `yaml:"providers"`
	Strategies struct {
		Fallback []StrategyConfig `yaml:"fallback"`
	} `yaml:"strategies"`
}

func LoadConfig(ctx context.Context, path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}

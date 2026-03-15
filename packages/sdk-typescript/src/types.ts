// Copyright (c) 2026 1Core Labs. MIT License.

export interface ChatRequest {
  prompt: string;
}

export interface ChatResponse {
  provider: string;
  model: string;
  completion: string;
  latency_ms?: number;
}

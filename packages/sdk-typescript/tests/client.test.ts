// Copyright (c) 2026 1Core Labs. MIT License.
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { AxonClient } from '../src/client';

// Mock the global fetch
global.fetch = vi.fn();

describe('AxonClient', () => {
  beforeEach(() => {
    vi.resetAllMocks();
  });

  it('should call fetch with correct parameters', async () => {
    const mockResponse = {
      ok: true,
      json: async () => ({
        provider: 'openai',
        model: 'gpt-4',
        completion: 'Hello from TS SDK mock',
        latency_ms: 15
      })
    };
    
    (global.fetch as any).mockResolvedValue(mockResponse);

    const client = new AxonClient({ apiKey: 'test-key', baseUrl: 'http://localhost:8080/v1' });
    const result = await client.complete('Hello');

    expect(global.fetch).toHaveBeenCalledWith('http://localhost:8080/v1/chat/completions', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer test-key'
      },
      body: JSON.stringify({ prompt: 'Hello' })
    });

    expect(result.completion).toBe('Hello from TS SDK mock');
  });
});

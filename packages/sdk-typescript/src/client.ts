// Copyright (c) 2026 1Core Labs. MIT License.
import { ChatRequest, ChatResponse } from './types';

export interface AxonClientOptions {
  apiKey?: string;
  baseUrl?: string;
}

export class AxonClient {
  private apiKey: string;
  private baseUrl: string;

  constructor(options: AxonClientOptions = {}) {
    this.apiKey = options.apiKey || process.env.AXON_API_KEY || '';
    this.baseUrl = options.baseUrl || process.env.AXON_BASE_URL || 'http://localhost:8080/v1';
  }

  async complete(prompt: string): Promise<ChatResponse> {
    const req: ChatRequest = { prompt };
    
    // Native fetch only, no external dependencies
    const response = await fetch(`${this.baseUrl}/chat/completions`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${this.apiKey}`
      },
      body: JSON.stringify(req)
    });

    if (!response.ok) {
      throw new Error(`Axon API error: ${response.status} ${response.statusText}`);
    }

    return await response.json() as ChatResponse;
  }
}

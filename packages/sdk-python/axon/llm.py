# Copyright (c) 2026 1Core Labs. MIT License.
import json
from typing import AsyncGenerator, Generator
from .client import AxonClient
from .types import ChatRequest, ChatResponse

class LLMModule:
    def __init__(self, client: AxonClient):
        self.client = client

    async def complete(self, prompt: str) -> ChatResponse:
        req = ChatRequest(prompt=prompt)
        resp = await self.client.post("/chat/completions", json=req.model_dump())
        resp.raise_for_status()
        return ChatResponse.model_validate(resp.json())

    def complete_sync(self, prompt: str) -> ChatResponse:
        req = ChatRequest(prompt=prompt)
        resp = self.client.post_sync("/chat/completions", json=req.model_dump())
        resp.raise_for_status()
        return ChatResponse.model_validate(resp.json())

    async def stream(self, prompt: str) -> AsyncGenerator[ChatResponse, None]:
        # Minimal mock implementation for stream (since Gateway doesn't support SSE yet, just yield once)
        req = ChatRequest(prompt=prompt)
        resp = await self.client.post("/chat/completions", json=req.model_dump())
        resp.raise_for_status()
        yield ChatResponse.model_validate(resp.json())

# Copyright (c) 2026 1Core Labs. MIT License.
from __future__ import annotations
import os
from contextlib import asynccontextmanager

from .client import AxonClient
from .llm import LLMModule

_client: AxonClient | None = None
llm: LLMModule | None = None

def init(api_key: str | None = None, base_url: str | None = None) -> None:
    global _client, llm
    if api_key is None:
        api_key = os.environ.get("AXON_API_KEY", "")
    if base_url is None:
        base_url = os.environ.get("AXON_BASE_URL", "http://localhost:8080/v1")
        
    _client = AxonClient(api_key=api_key, base_url=base_url)
    llm = LLMModule(_client)

@asynccontextmanager
async def async_init(api_key: str | None = None, base_url: str | None = None):
    init(api_key, base_url)
    try:
        yield
    finally:
        if _client:
            await _client.close()

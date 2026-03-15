# Copyright (c) 2026 1Core Labs. MIT License.
from pydantic import BaseModel, ConfigDict
from typing import Optional

class ChatRequest(BaseModel):
    model_config = ConfigDict(populate_by_name=True)
    prompt: str

class ChatResponse(BaseModel):
    model_config = ConfigDict(populate_by_name=True)
    provider: str
    model: str
    completion: str
    latency_ms: Optional[int] = None

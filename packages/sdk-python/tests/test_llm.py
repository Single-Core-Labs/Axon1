# Copyright (c) 2026 1Core Labs. MIT License.
import pytest
import respx
import httpx
from axon.client import AxonClient
from axon.llm import LLMModule
from axon.types import ChatResponse

@pytest.fixture
def client():
    return AxonClient(api_key="test-key", base_url="http://localhost:8080/v1")

@pytest.fixture
def llm(client):
    return LLMModule(client)

@pytest.mark.asyncio
@respx.mock
async def test_complete_async(llm):
    mock_resp = {
        "provider": "openai",
        "model": "gpt-4",
        "completion": "Hello from mock",
        "latency_ms": 10
    }
    
    respx.post("http://localhost:8080/v1/chat/completions").mock(
        return_value=httpx.Response(200, json=mock_resp)
    )

    resp = await llm.complete("Hello")
    assert isinstance(resp, ChatResponse)
    assert resp.provider == "openai"
    assert resp.completion == "Hello from mock"

@respx.mock
def test_complete_sync(llm):
    mock_resp = {
        "provider": "anthropic",
        "model": "claude",
        "completion": "Hi there sync",
        "latency_ms": 20
    }
    
    respx.post("http://localhost:8080/v1/chat/completions").mock(
        return_value=httpx.Response(200, json=mock_resp)
    )

    resp = llm.complete_sync("Hi")
    assert isinstance(resp, ChatResponse)
    assert resp.completion == "Hi there sync"

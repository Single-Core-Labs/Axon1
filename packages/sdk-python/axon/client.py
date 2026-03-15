# Copyright (c) 2026 1Core Labs. MIT License.
import httpx
from typing import Any

class AxonClient:
    def __init__(self, api_key: str, base_url: str):
        self.api_key = api_key
        self.base_url = base_url
        self.client = httpx.AsyncClient(
            base_url=self.base_url,
            headers={
                "Authorization": f"Bearer {self.api_key}",
                "Content-Type": "application/json",
            },
            timeout=30.0
        )
        self.sync_client = httpx.Client(
            base_url=self.base_url,
            headers={
                "Authorization": f"Bearer {self.api_key}",
                "Content-Type": "application/json",
            },
            timeout=30.0
        )

    async def post(self, url: str, json: Any) -> httpx.Response:
        return await self.client.post(url, json=json)
        
    def post_sync(self, url: str, json: Any) -> httpx.Response:
        return self.sync_client.post(url, json=json)

    async def close(self):
        await self.client.aclose()
        self.sync_client.close()

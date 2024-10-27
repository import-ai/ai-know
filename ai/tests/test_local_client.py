import json

import pytest

from app import v1_stream
from core.config import Config, load_config
from core.entity.api import ChatRequest
from core.pipeline import Pipeline
import json


@pytest.fixture(scope="function")
def pipeline() -> Pipeline:
    config: Config = load_config()
    return Pipeline(config)


@pytest.mark.parametrize("namespace, query", [
    ("test", "怎么安装 CUDA 驱动")
])
async def test_local_client(pipeline: Pipeline, namespace: str, query: str):
    request = ChatRequest(session_id="fake_id", query=query, namespace=namespace)
    async for each in v1_stream(pipeline, request):
        response = json.loads(each)
        if response["response_type"] == "delta":
            print(response["delta"], end="", flush=True)
        elif response["response_type"] == "citation":
            print("\n".join(["", "-" * 32, json.dumps(response["citation"])]))

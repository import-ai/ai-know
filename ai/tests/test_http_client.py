import tempfile
from typing import AsyncIterator

import pytest
import requests
from sseclient import SSEClient

from core.entity.api import ChatRequest, InsertRequest
from tests.test_local_client import assert_stream

namespace: str = "test"


@pytest.fixture(scope='session', autouse=True)
def init_before_all_tests():
    import os
    import subprocess
    import time

    with tempfile.TemporaryDirectory() as tmp_path:
        env = os.environ.copy()
        env["AI_KNOW_VECTOR_DB_PATH"] = tmp_path
        project_root = os.path.join(os.path.dirname(__file__), "..")
        process = subprocess.Popen(
            ["uvicorn", "--host", "127.0.0.1", "--port", "8000", "app:app"],
            cwd=project_root, env=env
        )

        while True:  # 等待服务起来
            try:
                response = requests.get("http://localhost:8000/api/v1/health")
                response.raise_for_status()
                break
            except requests.exceptions.ConnectionError:
                time.sleep(1)
        yield

        process.terminate()
        process.wait()


async def stream(reqeust: ChatRequest) -> AsyncIterator[str]:
    response = requests.post("http://localhost:8000/api/v1/stream", json=reqeust.model_dump(), stream=True)
    try:
        response.raise_for_status()
    except Exception as e:
        raise AssertionError(response.text) from e
    client = SSEClient(response)
    for event in client.events():
        yield event.data


def test_create_index():
    request = InsertRequest(title="下周计划", content="+ 9:00 起床\n+ 10:00 上班")
    response = requests.put(f"http://localhost:8000/api/v1/index/{namespace}/a", json=request.model_dump())
    response.raise_for_status()


@pytest.mark.parametrize("query", [
    "下周有什么计划"
])
async def test_stream(query: str):
    request = ChatRequest(session_id="fake_id", namespace=namespace, query=query)
    await assert_stream(stream(request))

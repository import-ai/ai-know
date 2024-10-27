import json
from typing import Iterator

import pytest
import requests
from sseclient import SSEClient

from core.entity.api import ChatRequest


@pytest.fixture(scope='session', autouse=True)
def init_before_all_tests():
    import os
    import subprocess
    import time

    env = os.environ.copy()
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


def stream(reqeust: ChatRequest) -> Iterator[dict]:
    response = requests.post("http://localhost:8000/api/v1/stream", json=reqeust.model_dump(), stream=True)
    try:
        response.raise_for_status()
    except Exception as e:
        raise AssertionError(response.text) from e
    client = SSEClient(response)
    for event in client.events():
        yield json.loads(event.data)


@pytest.mark.parametrize("namespace, query", [
    ("test", "如何安装 CUDA 驱动")
])
def test_stream(namespace: str, query: str):
    request = ChatRequest(session_id="fake_id", namespace=namespace, query=query)
    for response in stream(request):
        if response["response_type"] == "delta":
            print(response["delta"], end="", flush=True)
        elif response["response_type"] == "citation":
            print("\n".join(["", "-" * 32, json.dumps(response["citation"])]))

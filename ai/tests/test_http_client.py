import tempfile
from typing import AsyncIterator, List

import pytest
import requests
from sseclient import SSEClient

from core.entity.api import ChatRequest, InsertRequest
from test_local_client import query_test_case
from tests.test_local_client import assert_stream, create_test_case


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


@pytest.mark.parametrize(*create_test_case)
def test_create_index(namespace: str, element_id: str, title: str, content: str):
    request = InsertRequest(title=title, content=content)
    response = requests.put(f"http://localhost:8000/api/v1/index/{namespace}/{element_id}", json=request.model_dump())
    response.raise_for_status()


@pytest.mark.parametrize(*query_test_case)
async def test_stream(namespace: str, query: str, element_id_list: List[str]):
    request = ChatRequest(session_id="fake_id", namespace=namespace, query=query, element_id_list=element_id_list)
    await assert_stream(stream(request))

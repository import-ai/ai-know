import json
import pathlib
from typing import AsyncIterator, List

import pytest

from app import v1_stream
from core.config import Config, load_config
from core.entity.api import ChatRequest, InsertRequest
from core.entity.retrieve.chunk import Chunk
from core.ingestion import split_markdown
from core.pipeline import Pipeline

namespace: str = "test"


@pytest.fixture(scope="function")
def pipeline(tmp_path: pathlib.Path) -> Pipeline:
    config: Config = load_config()
    config.vector_db.path = str(tmp_path)
    return Pipeline(config)


async def assert_stream(stream: AsyncIterator[str]):
    async for each in stream:
        response = json.loads(each)
        if response["response_type"] == "delta":
            print(response["delta"], end="", flush=True)
        elif response["response_type"] == "citation_list":
            print("\n".join(["", "-" * 32, json.dumps(response["citation_list"])]))


async def test_index_create(pipeline: Pipeline):
    request = InsertRequest(title="下周计划", content="+ 9:00 起床\n+ 10:00 上班")
    chunk_list: List[Chunk] = split_markdown(namespace, "a", request.title, request.content)
    pipeline.retriever.vector_db.insert(chunk_list)


@pytest.mark.parametrize("query", [
    "下周有什么计划"
])
async def test_stream(pipeline: Pipeline, query: str):
    request = ChatRequest(session_id="fake_id", query=query, namespace=namespace)
    await assert_stream(v1_stream(pipeline, request))

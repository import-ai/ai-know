import json
import tempfile
from typing import AsyncIterator, List

import pytest

from app import v1_stream
from core.config import Config, load_config
from core.entity.api import ChatRequest
from core.entity.retrieve.chunk import Chunk
from core.ingestion import split_markdown
from core.pipeline import Pipeline


@pytest.fixture(scope="session")
def pipeline() -> Pipeline:
    with tempfile.TemporaryDirectory() as tmp_path:
        config: Config = load_config()
        config.vector_db.path = tmp_path
        yield Pipeline(config)


async def assert_stream(stream: AsyncIterator[str]):
    async for each in stream:
        response = json.loads(each)
        if response["response_type"] == "delta":
            print(response["delta"], end="", flush=True)
        elif response["response_type"] == "citation_list":
            print("\n".join(["", "-" * 32, json.dumps(response["citation_list"], ensure_ascii=False)]))


create_test_case = ("namespace, element_id, title, content", [
    ("ns_a", "e_id_a0", "下周计划", "+ 9:00 起床\n+ 10:00 上班"),
    ("ns_a", "e_id_a1", "下周计划", "+ 8:00 起床\n+ 9:00 上班"),
    ("ns_b", "e_id_b0", "下周计划", "+ 7:00 起床\n+ 8:00 上班"),
])


@pytest.mark.parametrize(*create_test_case)
async def test_index_create(pipeline: Pipeline, namespace: str, element_id: str, title: str, content: str):
    chunk_list: List[Chunk] = split_markdown(namespace, element_id, title, content)
    pipeline.retriever.vector_db.insert(chunk_list)


query_test_case = ("namespace, query, element_id_list", [
    ("ns_a", "下周有什么计划", ["e_id_a0"]),
    ("ns_a", "下周有什么计划", ["e_id_a1"]),
    ("ns_b", "下周有什么计划", ["e_id_b0"]),
])


@pytest.mark.parametrize(*query_test_case)
async def test_stream(pipeline: Pipeline, namespace: str, query: str, element_id_list: List[str]):
    assert pipeline.retriever.vector_db.collection.count() > 0
    request = ChatRequest(session_id="fake_id", query=query, namespace=namespace, element_id_list=element_id_list)
    await assert_stream(v1_stream(pipeline, request))

import pathlib
from typing import List, Tuple

import pytest

from core.entity.retrieve.chunk import Chunk, ChunkType
from core.retriever.vector_db import VectorDB

namespace = "pytest"

@pytest.fixture(scope="function")
def db(tmp_path: pathlib.Path) -> VectorDB:
    db = VectorDB(str(tmp_path), "BAAI/bge-m3", "cpu")
    chunk_list = [
        Chunk(element_id="a", text="apple", title="apple", chunk_type=ChunkType.keyword, namespace=namespace),
        Chunk(element_id="a", text="car", title="apple", chunk_type=ChunkType.keyword, namespace=namespace),
        Chunk(element_id="b", text="snake", title="snake", chunk_type=ChunkType.keyword, namespace=namespace)
    ]
    db.insert(chunk_list)
    yield db


@pytest.mark.parametrize("query, k, rank, expected_text, expected_element_id", [
    ("banana", 3, 0, "apple", "a"),
    ("bike", 3, 0, "car", "a"),
    ("chunk_type", 3, 0, "snake", "b")
])
def test_db_query(db: VectorDB, query: str, k: int, rank: int, expected_text: str, expected_element_id: str):
    assert db.collection.count() > 0
    result_list: List[Tuple[Chunk, float]] = db.query(namespace, query, k)
    assert len(result_list) == k
    assert result_list[rank][0].text == expected_text
    assert result_list[rank][0].element_id == expected_element_id


@pytest.mark.parametrize("element_id, expected_count", [("a", 1), ("b", 2)])
def test_db_remove(db: VectorDB, element_id: str, expected_count: int):
    assert db.collection.count() == 3
    db.remove(namespace, element_id)
    assert db.collection.count() == expected_count

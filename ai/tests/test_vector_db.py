import pathlib
from typing import List

import pytest

from core.retriever.embedding import VectorDB
from core.entity import Chunk, Retrieval, ChunkType


@pytest.fixture()
def db(tmp_path: pathlib.Path) -> VectorDB:
    db = VectorDB(str(tmp_path), "BAAI/bge-m3", "cpu")
    db.insert([
        Chunk(doc_id="a", text="apple", chunk_type=ChunkType.keyword),
        Chunk(doc_id="a", text="car", chunk_type=ChunkType.keyword),
        Chunk(doc_id="b", text="snake", chunk_type=ChunkType.keyword)
    ])
    yield db


@pytest.mark.parametrize("query, k, rank, expected_text, expected_doc_id", [
    ("banana", 3, 0, "apple", "a"),
    ("bike", 3, 0, "car", "a"),
    ("chunk_type", 3, 0, "snake", "b")
])
def test_db_query(db: VectorDB, query: str, k: int, rank: int, expected_text: str, expected_doc_id: str):
    result_list: List[Retrieval] = db.query(query, k)
    assert len(result_list) == k
    assert result_list[rank].chunk.text == expected_text
    assert result_list[rank].chunk.doc_id == expected_doc_id


@pytest.mark.parametrize("doc_id, expected_count", [("a", 1), ("b", 2)])
def test_db_remove(db: VectorDB, doc_id: str, expected_count: int):
    assert db.collection.count() == 3
    db.remove(doc_id)
    assert db.collection.count() == expected_count

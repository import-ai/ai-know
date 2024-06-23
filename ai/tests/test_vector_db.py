import pathlib
from typing import List

import pytest

from core.embedding import Embedding
from core.entity import Chunk, Retrieval, ChunkType


@pytest.fixture()
def db(tmp_path: pathlib.Path) -> Embedding:
    db = Embedding(str(tmp_path), "BAAI/bge-m3", "cpu")
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
def test_db(db: Embedding, query: str, k: int, rank: int, expected_text: str, expected_doc_id: str):
    result_list: List[Retrieval] = db.query(query, k)
    assert len(result_list) == k
    assert result_list[rank].chunk.text == expected_text
    assert result_list[rank].chunk.doc_id == expected_doc_id

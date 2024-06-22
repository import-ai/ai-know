import os
import pathlib
from typing import List

import pytest

from core.entity import Chunk, Retrieval
from core.vector import Recall


@pytest.fixture()
def db(tmp_path: pathlib.Path) -> Recall:
    db = Recall(os.path.join(tmp_path, "chroma.data"), "BAAI/bge-m3", "cpu")
    db.insert([Chunk(id="a", text="apple"), Chunk(id="b", text="car"), Chunk(id="c", text="snake")])
    yield db


@pytest.mark.parametrize("query, k, rank, expected_text", [
    ("banana", 3, 0, "apple"),
    ("bike", 3, 0, "car")
])
def test_db(db: Recall, query: str, k: int, rank: int, expected_text: str):
    result_list: List[Retrieval] = db.query(query, k)
    assert len(result_list) == k
    assert result_list[rank].chunk.text == expected_text

import os
import pathlib
from typing import List

import pytest

from entity import Chunk, Retrieval
from vector_db import VectorDB


@pytest.fixture()
def db(tmp_path: pathlib.Path) -> VectorDB:
    db = VectorDB(os.path.join(tmp_path, "chroma.data"))
    db.insert([Chunk(id="a", text="apple"), Chunk(id="b", text="car"), Chunk(id="c", text="snake")])
    yield db
    os.remove(db.datapath)


@pytest.mark.parametrize("query, k, rank, expected_text", [
    ("banana", 3, 0, "apple"),
    ("moto", 3, 1, "car")
])
def test_db(db: VectorDB, query: str, k: int, rank: int, expected_text: str):
    result_list: List[Retrieval] = db.query(query, k)
    assert len(result_list) == k
    assert result_list[rank][0].text == expected_text

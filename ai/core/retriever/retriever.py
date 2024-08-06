from .ranker import Ranker
from .vector_db import VectorDB
from typing import List, Tuple
from core.config import VectorDBConfig, RankerConfig
from core.entity import Retrieval


class Retriever:
    def __init__(self, vector_db_config: VectorDBConfig, ranker_config: RankerConfig):
        self.vector_db = VectorDB(**vector_db_config.model_dump())
        self.ranker = Ranker(**ranker_config.model_dump())

    def query(self, query: str, k: int = 3, threshold: float = 0) -> List[Retrieval]:
        retrieval_list: List[Retrieval] = self.vector_db.query(query, k << 2)
        rank_result: List[Tuple[int, float]] = self.ranker.rank(query, [r.chunk.text for r in retrieval_list])
        for i, score in rank_result:
            retrieval_list[i].score = score
        return [retrieval_list[i] for i, score in rank_result if score > threshold]

from typing import List, Tuple
from typing import Optional

from core.config import VectorDBConfig, RankerConfig
from core.entity.retrieve.chunk import Chunk, TextRetrieval
from core.entity.retrieve.retrieval import Score
from core.entity.trace_info import TraceInfo
from core.retriever.ranker import Ranker
from core.retriever.vector_db import VectorDB
from core.runner.base import BaseRunner


class Retriever(BaseRunner):
    def __init__(self, vector_db_config: VectorDBConfig, ranker_config: RankerConfig):
        self.vector_db = VectorDB(**vector_db_config.model_dump())
        self.ranker = Ranker(**ranker_config.model_dump())

    async def ainvoke(self, trace_info: TraceInfo, namespace: str = ..., query: str = ..., k: int = 3,
                      threshold: Optional[float] = None) -> List[TextRetrieval]:
        recall_result_list: List[Tuple[Chunk, float]] = self.vector_db.query(namespace, query, k << 2)
        rank_result: List[Tuple[int, float]] = self.ranker.rank(query, [c.text for c, _ in recall_result_list])

        retrieval_list: List[TextRetrieval] = [
            TextRetrieval(chunk=recall_result_list[i][0], score=Score(recall=recall_result_list[i][1], rerank=score))
            for i, score in rank_result
        ]
        return [r for r in retrieval_list if threshold is None or r.score.rerank > threshold]

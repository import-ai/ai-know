from typing import List, AsyncIterator

from core.config import Config
from core.entity.api import ChatRequest
from core.entity.retrieve.chunk import TextRetrieval
from core.entity.trace_info import TraceInfo
from core.rag import RAG
from core.retriever.retriever import Retriever
from core.runner.base import BaseRunner


class Pipeline(BaseRunner):

    def __init__(self, config: Config):
        self.retriever: Retriever = Retriever(config.vector_db, config.ranker)
        self.rag: RAG = RAG(config.openai)

    async def astream(self, trace_info: TraceInfo, request: ChatRequest = None, *args, **kwargs) -> AsyncIterator[dict]:
        query = request.query
        retrieval_list: List[TextRetrieval] = await self.retriever.ainvoke(trace_info, query, 3)
        async for delta in self.rag.astream(query, retrieval_list):
            yield {"response_type": "delta", "delta": delta}
        yield {"response_type": "citation", "citation": [r.to_reference() for r in retrieval_list]}

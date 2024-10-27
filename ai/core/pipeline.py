from typing import List, AsyncIterator

from core.config import Config
from core.entity.api import ChatRequest, ChatBaseResponse, ChatDeltaResponse, ChatCitationListResponse
from core.entity.retrieve.chunk import TextRetrieval
from core.entity.trace_info import TraceInfo
from core.rag import RAG
from core.retriever.retriever import Retriever
from core.runner.base import BaseRunner


class Pipeline(BaseRunner):

    def __init__(self, config: Config):
        self.retriever: Retriever = Retriever(config.vector_db, config.ranker)
        self.rag: RAG = RAG(config.openai)

    async def astream(self, trace_info: TraceInfo, request: ChatRequest = ...,
                      *args, **kwargs) -> AsyncIterator[ChatBaseResponse]:
        retrieval_list: List[TextRetrieval] = await self.retriever.ainvoke(
            trace_info, request.namespace, request.query, 3)
        async for delta in self.rag.astream(request.query, retrieval_list):
            yield ChatDeltaResponse(delta=delta)

        yield ChatCitationListResponse(citation_list=[r.to_citation() for r in retrieval_list])

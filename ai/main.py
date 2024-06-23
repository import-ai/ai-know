from contextlib import asynccontextmanager
import json
from functools import partial
from typing import List, Iterator
from sse_starlette import EventSourceResponse

import uvicorn
from fastapi import FastAPI, Request, status
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse, Response
from pydantic import BaseModel, Field

from core.config import load_config, Config
from core.embedding import Embedding
from core.entity import Chunk, Retrieval
from core.ingestion import split_markdown
from core.logger import get_logger
from core.rag import RAG


class InsertRequest(BaseModel):
    title: str = Field(description="Document title")
    content: str = Field(description="Document content")


dumps = partial(json.dumps, ensure_ascii=False, separators=(",", ":"))

config: Config = load_config()
logger = get_logger(__name__)
embedding: Embedding = ...
rag: RAG = ...


def init():
    global embedding, rag
    embedding = Embedding("chroma_data", config.embedding_model_name_or_path, config.device)
    rag = RAG()


@asynccontextmanager
async def lifespan(_: FastAPI):
    init()
    yield
    global embedding, rag
    del embedding, rag


app = FastAPI(lifespan=lifespan, title="AI Knowledge", version="0.0.1", description="")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"]
)


@app.exception_handler(Exception)
async def exception_handler(_: Request, e: Exception) -> Response:
    return JSONResponse(
        status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
        content={
            "error": f"{e.__class__.__name__}: {str(e)}",
            "status": 500
        }
    )


@app.put("/api/index/{doc_id}", response_model=None)
async def create_or_update(doc_id: str, request: InsertRequest):
    chunk_list: List[Chunk] = split_markdown(doc_id, request.title, request.content)
    embedding.insert(chunk_list)


@app.get("/api/index/recall", response_model=List[Retrieval])
async def recall(query: str, k: int):
    retrieval_list: List[Retrieval] = embedding.query(query, k)
    return retrieval_list


def stream_chat(query: str) -> Iterator[str]:
    retrieval_list: List[Retrieval] = embedding.query(query, 30)
    for delta in rag.chat(query, retrieval_list):  # TODO Add rerank
        yield dumps({"response_type": "delta", "content": delta})
    yield dumps({"response_type": "citation", "content": [r.chunk.model_dump() for r in retrieval_list]})
    yield '[DONE]'


@app.post("/api/chat")
async def api_chat(query: str):
    return EventSourceResponse(stream_chat(query))


# healthcheck
@app.get("/api/health")
async def healthcheck():
    return {"status": 200}


def main():
    uvicorn.run(
        'main:app',
        host='0.0.0.0',
        port=config.port,
        workers=config.workers
    )


if __name__ == '__main__':
    main()

from contextlib import asynccontextmanager
import json
from datetime import datetime
from functools import partial
from typing import List, Iterator
from sse_starlette import EventSourceResponse

import uvicorn
from fastapi import FastAPI, Request, status
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse, Response, HTMLResponse
from pydantic import BaseModel, Field

from core.config import load_config, Config
from core.entity import Chunk, Retrieval
from core.ingestion import split_markdown
from core.logger import get_logger
from core.rag import RAG
from core.retriever.retriever import Retriever

dumps = partial(json.dumps, ensure_ascii=False, separators=(",", ":"))

with open("resource/index.html") as f:
    chat_html: str = f.read()
start_time: datetime = datetime.now()
config: Config = load_config()
logger = get_logger(__name__)
retriever: Retriever = ...
rag: RAG = ...


class InsertRequest(BaseModel):
    title: str = Field(description="Document title")
    content: str = Field(description="Document content")


def init():
    global retriever, rag
    retriever = Retriever(config.vector_db, config.ranker)
    rag = RAG(config.openai)


@asynccontextmanager
async def lifespan(_: FastAPI):
    init()
    yield
    global retriever, rag
    del retriever, rag


app = FastAPI(lifespan=lifespan, title="AI Know", description="")

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
    retriever.vector_db.remove(doc_id)
    chunk_list: List[Chunk] = split_markdown(doc_id, request.title, request.content)
    retriever.vector_db.insert(chunk_list)


@app.get("/api/index/recall", response_model=List[Retrieval])
async def recall(query: str, k: int):
    retrieval_list: List[Retrieval] = retriever.query(query, k)
    return retrieval_list


def stream_chat(query: str) -> Iterator[str]:
    try:
        retrieval_list: List[Retrieval] = retriever.query(query, 3)

        doc_id_list = []
        for retrieval in retrieval_list:
            if retrieval.chunk.doc_id not in doc_id_list:
                doc_id_list.append(retrieval.chunk.doc_id)

        for delta in rag.chat(query, retrieval_list):
            yield dumps({"response_type": "delta", "content": delta})
        yield dumps({"response_type": "citation", "content": [r.chunk.model_dump() for r in retrieval_list]})
    except Exception as e:
        yield dumps({
            "response_type": "error",
            "exception_class": e.__class__.__name__,
            "exception_message": str(e)
        })
    yield '[DONE]'


@app.get("/api/chat")
async def api_chat(query: str):
    return EventSourceResponse(stream_chat(query))


@app.get("/", response_class=HTMLResponse)
async def index():
    return HTMLResponse(content=chat_html, status_code=200)


# healthcheck
@app.get("/api/health")
async def healthcheck():
    return {
        "status": 200,
        "uptime": str(datetime.now() - start_time)
    }


def main():
    uvicorn.run(
        'main:app',
        host='0.0.0.0',
        port=config.port,
        workers=config.workers
    )


if __name__ == '__main__':
    main()

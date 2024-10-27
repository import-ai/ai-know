import json
from contextlib import asynccontextmanager
from datetime import datetime
from functools import partial
from typing import List, AsyncIterator

from fastapi import FastAPI, Request, status, APIRouter
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse, Response
from fastapi.responses import StreamingResponse
from pydantic import BaseModel, Field

from core.config import load_config, Config
from core.entity.api import ChatRequest
from core.entity.retrieve.chunk import Chunk
from core.entity.trace_info import TraceInfo
from core.ingestion import split_markdown
from core.pipeline import Pipeline
from core.util.logger import get_logger

dumps = partial(json.dumps, ensure_ascii=False, separators=(",", ":"))

start_time: datetime = datetime.now()
config: Config = load_config()
logger = get_logger(__name__)
pipeline: Pipeline


class InsertRequest(BaseModel):
    title: str = Field(description="Document title")
    content: str = Field(description="Document content")


def init():
    global pipeline
    pipeline = Pipeline(config)


@asynccontextmanager
async def lifespan(_: FastAPI):
    init()
    yield


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


v1 = APIRouter(prefix="/api/v1")


@v1.put("/index/{namespace}/{element_id}", response_model=None)
async def create_or_update(namespace: str, element_id: str, request: InsertRequest):
    pipeline.retriever.vector_db.remove(namespace, element_id)
    chunk_list: List[Chunk] = split_markdown(namespace, element_id, request.title, request.content)
    for chunk in chunk_list:
        chunk.namespace = namespace
    pipeline.retriever.vector_db.insert(chunk_list)


@v1.delete("/index/{namespace}/{element_id}", response_model=None)
async def delete(namespace: str, element_id: str):
    pipeline.retriever.vector_db.remove(namespace, element_id)


async def v1_stream(p: Pipeline, request: ChatRequest) -> AsyncIterator[str]:
    trace_info = TraceInfo(logger=logger.getChild(request.session_id))
    try:
        async for delta in p.astream(trace_info, request):
            yield dumps(delta)
    except Exception as e:
        yield dumps({
            "response_type": "error",
            "message": "Unknown error"
        })
        trace_info.logger.exception({
            "exception_class": e.__class__.__name__,
            "exception_message": str(e)
        })
    yield dumps({"response_type": "done"})


@v1.post("/stream")
async def api_v1_stream(request: ChatRequest):
    return StreamingResponse(v1_stream(pipeline, request), media_type="text/event-stream")


# healthcheck
@v1.get("/health")
async def api_v1_health():
    return {
        "status": 200,
        "uptime": str(datetime.now() - start_time)
    }


app.include_router(v1)

import json
from contextlib import asynccontextmanager
from datetime import datetime
from functools import partial
from typing import List, AsyncIterator, Union

from fastapi import FastAPI, Request, status, APIRouter
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse, Response
from sse_starlette import EventSourceResponse

from core.config import load_config, Config
from core.entity.api import ChatRequest, InsertRequest, ChatDeltaResponse, ChatCitationListResponse, ChatBaseResponse
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


def init():
    global pipeline
    pipeline = Pipeline(config)


@asynccontextmanager
async def lifespan(_: FastAPI):
    # init()
    yield


app = FastAPI(lifespan=lifespan, title="AI Know", description="""## Description

+ **Index**: user's element, like text, url, photo, etc.
+ **LLM**: QA

## Pipeline

1. Create or update index
2. Answer the question""")

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


@v1.put("/index/{namespace}/{element_id}", response_model=None, tags=["Index"])
async def create_or_update_index(namespace: str, element_id: str, request: InsertRequest):
    pipeline.retriever.vector_db.remove(namespace, element_id)
    chunk_list: List[Chunk] = split_markdown(namespace, element_id, request.title, request.content)
    pipeline.retriever.vector_db.insert(chunk_list)


@v1.delete("/index/{namespace}/{element_id}", response_model=None, tags=["Index"])
async def delete_index(namespace: str, element_id: str):
    pipeline.retriever.vector_db.remove(namespace, element_id)


async def v1_stream(p: Pipeline, request: ChatRequest) -> AsyncIterator[str]:
    trace_info = TraceInfo(logger=logger.getChild(request.session_id))
    try:
        async for delta in p.astream(trace_info, request):
            yield dumps(delta.model_dump())
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


@v1.post("/stream", tags=["LLM"], response_model=Union[ChatBaseResponse, ChatDeltaResponse, ChatCitationListResponse])
async def api_v1_stream(request: ChatRequest):
    """
    Answer the query based on user's database.
    """
    return EventSourceResponse(v1_stream(pipeline, request))


# healthcheck
@v1.get("/health", tags=["Metrics"])
async def api_v1_health():
    return {
        "status": 200,
        "uptime": str(datetime.now() - start_time)
    }


app.include_router(v1)

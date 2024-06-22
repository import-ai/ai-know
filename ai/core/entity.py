from pydantic import BaseModel, Field
from typing import Optional
import shortuuid
import time
from enum import Enum


class ChunkType(str, Enum):
    doc: str = "doc"  # Whole document
    section: str = "section"  # Part of document
    snippet: str = "snippet"  # Part of section
    keyword: str = "keyword"


class Chunk(BaseModel):
    doc_id: str = Field(description="ID of note")
    text: str = Field(description="Chunk content")
    chunk_type: ChunkType = Field(description="Chunk type")
    start_lineno: int = Field(description="The start line number of this chunk, line included")
    end_lineno: int = Field(description="The end line number of this chunk, line excluded")

    chunk_id: str = Field(description="ID of chunk", default_factory=shortuuid.uuid)
    created_timestamp: float = Field(description="Unix timestamp in float format", default_factory=time.time)

    parent_chunk_id: Optional[str] = Field(description="A chunk could be split into many smaller chunks", default=None)


class Retrieval(BaseModel):
    chunk: Chunk
    distance: float

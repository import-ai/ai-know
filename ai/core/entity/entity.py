from pydantic import BaseModel, Field


class Chunk(BaseModel):
    id: str = Field(description="ID of note")
    text: str = Field(description="Chunk content")


class Retrieval(BaseModel):
    chunk: Chunk
    distance: float


__all__ = ["Chunk", "Retrieval"]

from typing import List, Tuple

from pydantic import BaseModel, Field


class Chunk(BaseModel):
    id: str = Field(description="ID of note")
    text: str = Field(description="Chunk content")


Retrieval = Tuple[Chunk, float]

from pydantic import BaseModel, Field


class Retrieval(BaseModel):
    id: int = Field(description="ID of note")
    content: str = Field(description="Full content")

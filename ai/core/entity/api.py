from typing import List, Optional

from pydantic import BaseModel, Field


class InsertRequest(BaseModel):
    title: str = Field(description="Document title")
    content: str = Field(description="Document content")


class ChatRequest(BaseModel):
    session_id: str
    query: str
    namespace: str
    element_id_list: Optional[List[str]] = Field(default=None)

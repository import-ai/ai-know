from enum import Enum

from pydantic import BaseModel, Field


class SelectorType(str, Enum):
    tag = "TAG"
    folder = "FOLDER"
    element = "ELEMENT"


class Selector(BaseModel):
    type: SelectorType
    id: str


class ChatRequest(BaseModel):
    session_id: str
    query: str
    namespace: str
    selector: Selector = Field(default=None)

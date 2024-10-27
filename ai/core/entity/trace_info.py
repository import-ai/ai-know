from logging import Logger

from pydantic import BaseModel


class TraceInfo(BaseModel):
    model_config = {"arbitrary_types_allowed": True}
    logger: Logger

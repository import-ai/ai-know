from typing import Any, AsyncIterator

from core.entity.trace_info import TraceInfo


class BaseRunner:
    async def astream(self, trace_info: TraceInfo, /, *args, **kwargs) -> Any:
        raise NotImplementedError

    async def ainvoke(self, trace_info: TraceInfo, /, *args, **kwargs) -> AsyncIterator[Any]:
        raise NotImplementedError

from datetime import datetime
from functools import partial
from typing import List, AsyncIterator

from openai import AsyncOpenAI

from core.config import OpenAIConfig
from core.entity.retrieve.retrieval import BaseRetrieval
from core.util.resource_util import resource


class RAG:

    def __init__(self, config: OpenAIConfig):
        self.client = AsyncOpenAI(api_key=config.api_key, base_url=config.base_url)
        self._chat = partial(self.client.chat.completions.create, model=config.model)
        with resource.open("prompt.md") as f:
            self.prompt: str = f.read()

    @classmethod
    def build_context(cls, retrieval_list: List[BaseRetrieval]) -> str:
        retrieval_prompt_list: List[str] = []
        for i, retrieval in enumerate(retrieval_list):
            prompt_list: List[str] = [
                f"[[{i + 1}]]",
                f"Snippet: {retrieval.to_prompt()}"
            ]
            retrieval_prompt_list.append("\n".join(prompt_list))
        return "\n\n".join(retrieval_prompt_list)

    def messages_prepare(self, query: str, template: str, retrieval_list: List[BaseRetrieval]) -> List[dict]:
        context = self.build_context(retrieval_list)
        messages = [
            {"role": "system", "content": template.format_map({
                "datetime": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
                "context": context
            })},
            {"role": "user", "content": query}
        ]
        return messages

    async def astream(self, query: str, retrieval_list: List[BaseRetrieval]) -> AsyncIterator[str]:
        messages = self.messages_prepare(query, self.prompt, retrieval_list)
        async for response in await self._chat(messages=messages, stream=True):
            if delta := response.choices[0].delta.content:
                yield delta

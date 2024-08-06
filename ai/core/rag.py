from datetime import datetime
from functools import partial
from typing import List

from openai import OpenAI

from core.config import OpenAIConfig
from core.entity import Retrieval, Chunk


class RAG:

    def __init__(self, config: OpenAIConfig):
        self.client = OpenAI(api_key=config.api_key, base_url=config.base_url)
        self._chat = partial(self.client.chat.completions.create, model=config.model)
        with open("resource/prompt.md") as f:
            self.prompt: str = f.read()

    @classmethod
    def build_context(cls, retrieval_list: List[Retrieval]) -> str:
        retrival_prompt_list: List[str] = []
        for i, retrieval in enumerate(retrieval_list):
            chunk: Chunk = retrieval.chunk
            prompt_list: List[str] = [
                f"[[{i + 1}]]",
                f"Snippet: {chunk.text}"
            ]
            retrival_prompt_list.append("\n".join(prompt_list))
        return "\n\n".join(retrival_prompt_list)

    def messages_prepare(self, query: str, template: str, retrieval_list: List[Retrieval]) -> List[dict]:
        context = self.build_context(retrieval_list)
        messages = [
            {"role": "system", "content": template.format_map({
                "datetime": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
                "context": context
            })},
            {"role": "user", "content": query}
        ]
        return messages

    def chat(self, query: str, retrieval_list: List[Retrieval]):
        for response in self._chat(messages=self.messages_prepare(query, self.prompt, retrieval_list), stream=True):
            if delta := response.choices[0].delta.content:
                yield delta

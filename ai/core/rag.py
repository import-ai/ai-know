from functools import partial
from typing import List, Dict, Any

from openai import OpenAI

from core.config import config
from core.entity import Retrieval, Chunk


class RAG:

    def __init__(self):
        self.client = OpenAI(api_key=config.openai_api_key, base_url=config.openai_base_url)
        self._chat = partial(self.client.chat.completions.create, model=config.openai_model)
        with open('resource/prompt.txt') as f:
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
            {"role": "system", "content": template.format_map({"context": context})},
            {"role": "user", "content": query}
        ]
        return messages

    def chat(self, query: str, retrieval_list: List[Retrieval]):
        for response in self._chat(messages=self.messages_prepare(query, self.prompt, retrieval_list), stream=True):
            if delta := response.choices[0].delta.content:
                yield delta

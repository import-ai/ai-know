import os
from functools import partial
from typing import List, Dict, Any
from urllib.parse import urljoin

import requests
from openai import OpenAI

from core.entity import Chunk as Retrieval


class Retriever:
    def __init__(self, base_url: str):
        self.base_url: str = base_url

    def retrieve(self, query: str) -> List[Retrieval]:
        retrieval_list: List[Retrieval] = []
        response: requests.Response = requests.post(urljoin(self.base_url, "/api/get_all_notes"), json={})
        response.raise_for_status()
        json_response: dict = response.json()
        for note in json_response["notes"]:
            retrieval_list.append(Retrieval.model_validate(note))
        return retrieval_list


class RAG:

    def __init__(self, retriever: Retriever):
        self.retriever: Retriever = retriever
        self.client = OpenAI(base_url=os.environ["OPENAI_BASE_URL"], api_key=os.environ["OPENAI_API_KEY"])
        self.chat = partial(self.client.chat.completions.create, model=os.environ["OPENAI_MODEL_NAME"])
        with open('../resource/prompt.txt') as f:
            self.prompt: str = f.read()

    @classmethod
    def build_context(cls, retrieval_list: List[Retrieval]) -> str:
        retrival_prompt_list: List[str] = []
        for retrieval in retrieval_list:
            prompt_list: List[str] = [
                f"[[{retrieval.doc_id}]]",
                f"Content: {retrieval.text}",
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

    def search(self, query: str):
        retrieval_list: List[Retrieval] = self.retriever.retrieve(query)
        assert len(retrieval_list) > 0, "Empty retrieval result"

        citations: List[Dict[str, Any]] = [r.model_dump() for r in retrieval_list]

        for response in self.chat(messages=self.messages_prepare(query, self.prompt, retrieval_list), stream=True):
            if delta := response.choices[0].delta.content:
                print(delta, end="")
        print("=" * 16)
        print(citations)


def main():
    retriever = Retriever(os.environ["BASE_URL"])
    rag = RAG(retriever)
    rag.search("Debian 的镜像")


if __name__ == '__main__':
    main()

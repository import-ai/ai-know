from typing import List

import chromadb
from chromadb.utils.embedding_functions import SentenceTransformerEmbeddingFunction

from core.entity import Chunk, Retrieval


class Embedding:
    def __init__(self, datapath: str, model_name: str = "BAAI/bge-m3", device: str = "cpu"):
        self.datapath: str = datapath
        self.client = chromadb.PersistentClient(path=datapath)
        self.collection = self.client.get_or_create_collection(
            name="default", metadata={"hnsw:space": "cosine"},
            embedding_function=SentenceTransformerEmbeddingFunction(model_name=model_name, device=device)
        )

    def insert(self, chunk_list: List[Chunk]):
        self.collection.add(
            documents=[c.text for c in chunk_list],
            ids=[c.chunk_id for c in chunk_list],
            metadatas=[c.metadata for c in chunk_list]
        )

    def query(self, query: str, k: int) -> List[Retrieval]:
        batch_result_list: chromadb.QueryResult = self.collection.query(query_texts=[query], n_results=k)
        result_list: List[Retrieval] = []
        for chunk_id, document, metadata, distance in zip(
                batch_result_list["ids"][0],
                batch_result_list["documents"][0],
                batch_result_list["metadatas"][0],
                batch_result_list["distances"][0],
        ):
            result_list.append(Retrieval(chunk=Chunk(chunk_id=chunk_id, text=document, **metadata), distance=distance))
        return result_list
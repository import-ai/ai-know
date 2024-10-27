from typing import List, Tuple

import chromadb
from chromadb.utils.embedding_functions import sentence_transformer_embedding_function

from core.entity.retrieve.chunk import Chunk

sef = sentence_transformer_embedding_function.SentenceTransformerEmbeddingFunction


class VectorDB:
    def __init__(self, path: str, model_name_or_path: str = "BAAI/bge-m3", device: str = "cpu", batch_size: int = 1):
        self.data_dir: str = path
        self.client = chromadb.PersistentClient(path=path)
        self.collection = self.client.get_or_create_collection(
            name="default", metadata={"hnsw:space": "cosine"},
            embedding_function=sef(model_name=model_name_or_path, device=device)
        )
        self.batch_size: int = batch_size

    def insert(self, chunk_list: List[Chunk]):
        for i in range(0, len(chunk_list), self.batch_size):
            batch: List[Chunk] = chunk_list[i:i + self.batch_size]
            self.collection.add(
                documents=[c.text for c in batch],
                ids=[c.chunk_id for c in batch],
                metadatas=[c.metadata for c in batch]
            )

    def remove(self, namespace: str, element_id: str):
        self.collection.delete(where={"$and": [{"namespace": namespace}, {"element_id": element_id}]})

    def query(self, namespace: str, query: str, k: int, element_id_list: List[str] = None) -> List[Tuple[Chunk, float]]:
        where = {"namespace": namespace}
        if element_id_list is not None:
            where = {"$and": [where, {"element_id": {"$in": element_id_list}}]}

        batch_result_list: chromadb.QueryResult = self.collection.query(query_texts=[query], n_results=k, where=where)
        result_list: List[Tuple[Chunk, float]] = []
        for chunk_id, document, metadata, distance in zip(
                batch_result_list["ids"][0],
                batch_result_list["documents"][0],
                batch_result_list["metadatas"][0],
                batch_result_list["distances"][0],
        ):
            result_list.append((Chunk(chunk_id=chunk_id, text=document, **metadata), distance))
        return result_list

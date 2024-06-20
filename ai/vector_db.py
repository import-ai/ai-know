from typing import List

import chromadb
from chromadb.utils.embedding_functions import ONNXMiniLM_L6_V2

from entity import Chunk, Retrieval


class VectorDB:
    def __init__(self, datapath: str):
        self.datapath: str = datapath
        self.client = chromadb.PersistentClient(path=datapath)
        self.collection = self.client.get_or_create_collection(
            name="default", metadata={"hnsw:space": "cosine"}, embedding_function=ONNXMiniLM_L6_V2())

    def insert(self, chunk_list: List[Chunk]):
        self.collection.add(documents=[c.text for c in chunk_list], ids=[c.id for c in chunk_list])

    def query(self, query: str, k: int) -> List[Retrieval]:
        batch_result_list: chromadb.QueryResult = self.collection.query(query_texts=[query], n_results=k)
        result_list: List[Retrieval] = []
        for idx, distance, document in zip(
                batch_result_list["ids"][0], batch_result_list["distances"][0], batch_result_list["documents"][0]):
            result_list.append((Chunk(id=idx, text=document), distance))
        return result_list

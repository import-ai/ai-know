from .ranker import Ranker
from .vector_db import VectorDB
from typing import List
from core.config import Config
from core.entity import Retrieval


class Retriever:
    def __init__(self, config: Config):
        self.vector_db = VectorDB(
            config.data_dir, config.vector_db_model_name_or_path, config.vector_db_device, config.vector_db_batch_size)
        self.ranker = Ranker(config.ranker_model_name_or_path, config.ranker_device, config.ranker_batch_size)

    def query(self, query: str) -> List[Retrieval]:
        pass


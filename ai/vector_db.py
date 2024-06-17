import os
from typing import List, Tuple

import duckdb
import pandas as pd

DATAPATH: str = "data.duckdb"


class VectorDB:
    def __init__(self, datapath: str):
        self.datapath: str = datapath
        if not os.path.exists(datapath):
            with open("resource/init.sql") as f:
                self.execute(f.read())

    def execute(self, query, parameters: Tuple = None, multiple_parameter_sets: bool = False):
        with duckdb.connect(self.datapath) as connection:
            connection.install_extension("vss")
            connection.load_extension("vss")
            return connection.execute(query, parameters, multiple_parameter_sets).df()

    def insert(self, doc_id, vector: List[float]):
        return self.execute("INSERT INTO vectors VALUES (?, ?);", (doc_id, vector))

    def query(self, vector: List[float], k: int) -> pd.DataFrame:
        return self.execute("""
            SELECT doc_id, vec, array_cosine_similarity(vec, CAST(? AS VECTOR)) AS similarity
            FROM vectors
            ORDER BY similarity DESC
            LIMIT ?;
        """, (vector, k))

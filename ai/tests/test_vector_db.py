import unittest
from vector_db import VectorDB
import os


class TestVectorDB(unittest.TestCase):
    test_db: str = "test.duckdb"

    @classmethod
    def clear_temp(cls):
        if os.path.exists(cls.test_db):
            os.remove(cls.test_db)

    @classmethod
    def setUpClass(cls):
        cls.clear_temp()
        db = VectorDB(cls.test_db)
        db.insert(1, [0.1, 0.2, 0.3])
        db.insert(2, [0.4, 0.5, 0.6])
        db.insert(3, [0.7, 0.8, 0.9])

    def setUp(self):
        self.db = VectorDB(self.test_db)

    def test_query(self):
        result = self.db.query([0.1, 0.2, 0.3], 2)
        self.assertEqual(len(result), 2)
        self.assertEqual(result.iloc[0]["doc_id"], 1)
        self.assertEqual(result.iloc[0]["similarity"], 1.0)
        self.assertEqual(result.iloc[1]["doc_id"], 2)
        self.assertGreater(result.iloc[1]["similarity"], 0.9)

    @classmethod
    def tearDownClass(cls):
        cls.clear_temp()

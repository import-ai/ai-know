from core.retriever.ranker import Ranker
from core.config import config


def test_common_ranker():
    ranker = Ranker(config.ranker.model_name_or_path, config.ranker.device, config.ranker.batch_size)
    candidate_list = ["hi", "panda is panda"]
    sorted_idx_list = ranker.rank("What's panda", candidate_list)
    assert sorted_idx_list[0] == 1

from core.ranker import CommonRanker
from core.config import config


def test_common_ranker():
    ranker = CommonRanker(config.ranker_model_name_or_path, config.ranker_device)
    candidate_list = ["hi", "panda is panda"]
    sorted_idx_list = ranker.rank("What's panda", candidate_list)
    assert sorted_idx_list[0] == 1

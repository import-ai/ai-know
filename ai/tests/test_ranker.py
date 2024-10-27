from core.config import Config, load_config
from core.retriever.ranker import Ranker


def test_common_ranker():
    config: Config = load_config()
    ranker = Ranker(config.ranker.model_name_or_path, config.ranker.device, config.ranker.batch_size)
    candidate_list = ["hi", "panda is panda"]
    sorted_idx_list = ranker.rank("What's panda", candidate_list)
    assert sorted_idx_list[0][0] == 1

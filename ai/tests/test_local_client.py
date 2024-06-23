from main import init, stream_chat
import json
import pytest


@pytest.mark.skip
def test_local_client():
    init()
    for each in stream_chat("怎么安装 CUDA 驱动"):
        response = json.loads(each)
        if response["response_type"] == "delta":
            print(response["content"], end="", flush=True)
        elif response["response_type"] == "citation":
            print(response["content"])

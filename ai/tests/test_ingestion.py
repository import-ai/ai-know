import pytest

from core.ingestion import split_markdown


def test_ingestion():
    markdown_text = """First part
# Level 1
## Level 2
Some text
## Level 2
More text
# Level 1
## Level 2
Last part"""

    chunk_list = split_markdown("asdf", markdown_text)
    assert len(chunk_list) == 7
    for i in [0, 1, 2, 5]:
        assert chunk_list[i].parent_chunk_id is None
    for i in [3, 4]:
        assert chunk_list[i].parent_chunk_id == chunk_list[2].chunk_id
    for i in [6]:
        assert chunk_list[i].parent_chunk_id == chunk_list[5].chunk_id

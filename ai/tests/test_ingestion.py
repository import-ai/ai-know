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

    chunk_list = split_markdown("id", "title", markdown_text)
    offset: int = 2
    assert len(chunk_list) == offset + 6
    for i in range(offset + 1):
        assert chunk_list[i].parent_chunk_id is None
    for i in [1, 4]:
        assert chunk_list[i + offset].parent_chunk_id is None
    for i in [2, 3]:
        assert chunk_list[i + offset].parent_chunk_id == chunk_list[1 + offset].chunk_id
    for i in [5]:
        assert chunk_list[i + offset].parent_chunk_id == chunk_list[4 + offset].chunk_id

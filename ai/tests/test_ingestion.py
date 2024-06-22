import pytest

from core.ingestion import split_markdown

markdown_text = """First part
# Level 1
## Level 2
Some text
## Level 2
More text
# Level 1
## Level 2
Last part"""


def test_ingestion():
    chunk_list = split_markdown("asdf", markdown_text)
    assert len(chunk_list) == 7


@pytest.mark.skip
def test_llama_index():
    from llama_index.core.node_parser import MarkdownNodeParser
    from llama_index.core import Document

    parser = MarkdownNodeParser()
    nodes = parser.get_nodes_from_documents([Document(text=markdown_text)])
    print(nodes)

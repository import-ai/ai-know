from llama_index.core.node_parser import MarkdownNodeParser
from core.entity import Chunk, ChunkType
from llama_index.core import Document

from typing import List


def split_markdown(doc_id: str, markdown_text: str) -> List[Chunk]:
    chunk_list: List[Chunk] = [Chunk(doc_id=doc_id, text=markdown_text, chunk_type=ChunkType.doc)]

    doc = Document(id_=doc_id, text=markdown_text)

    parser = MarkdownNodeParser()
    nodes = parser.get_nodes_from_documents([doc])

    for node in nodes:
        chunk_type = ChunkType.section if node.get_metadata_str() else ChunkType.snippet
        chunk: Chunk = Chunk(doc_id=doc_id, text=node.get_content(), chunk_type=chunk_type)

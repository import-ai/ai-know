from core.entity import Chunk, ChunkType
from typing import List
from pydantic import ValidationError


def line_level(line: str) -> int:
    return len(line) - len(line.lstrip('#'))


def split_markdown(doc_id: str, markdown_text: str) -> List[Chunk]:
    lines: List[str] = markdown_text.split('\n')
    chunk_list: List[Chunk] = [
        Chunk(doc_id=doc_id, text=markdown_text, chunk_type=ChunkType.doc, start_lineno=0, end_lineno=len(lines))
    ]
    header_lineno_stack: List[int] = []
    chunk_stack: List[Chunk] = []
    for lineno, line in enumerate(lines + ["#"]):  # Add a "#" to make program handle last part without special check.
        if line.startswith('#'):
            if len(header_lineno_stack) == 0:
                if lineno > 0 and '\n'.join(lines[:lineno]).strip():
                    chunk_list.append(Chunk(
                        doc_id=doc_id,
                        text='\n'.join(lines[:lineno]).strip(),
                        chunk_type=ChunkType.snippet,
                        start_lineno=0,
                        end_lineno=lineno
                    ))
                header_lineno_stack.append(lineno)
                continue
            start_lineno = header_lineno_stack[-1]
            current_level = line_level(line)
            while len(header_lineno_stack) > 0 and line_level(lines[header_lineno_stack[-1]]) >= current_level:
                header_lineno_stack.pop()
            header_lineno_stack.append(lineno)

            chunk = Chunk(
                doc_id=doc_id,
                text='\n'.join(lines[start_lineno:lineno]).strip(),
                chunk_type=ChunkType.section,
                start_lineno=start_lineno,
                end_lineno=lineno
            )

            while len(chunk_stack) > 0 and line_level(lines[chunk_stack[-1].start_lineno]) >= current_level:
                chunk_stack.pop()

            if len(chunk_stack) > 0:
                chunk.parent_chunk_id = chunk_stack[-1].chunk_id
            chunk_stack.append(chunk)
            chunk_list.append(chunk)

    return chunk_list


"""
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


"""

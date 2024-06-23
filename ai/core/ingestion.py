from core.entity import Chunk, ChunkType
from typing import List
from pydantic import ValidationError


def line_level(line: str) -> int:
    return len(line) - len(line.lstrip('#'))


def split_markdown(doc_id: str, title: str, markdown_text: str) -> List[Chunk]:
    lines: List[str] = markdown_text.split('\n')
    title = Chunk(doc_id=doc_id, text=title, chunk_type=ChunkType.title)
    body = Chunk(doc_id=doc_id, text=markdown_text, chunk_type=ChunkType.doc, start_lineno=0, end_lineno=len(lines))
    chunk_list: List[Chunk] = [title, body]
    chunk_stack: List[Chunk] = []
    previous_lineno: int = 0
    for lineno, line in enumerate(lines + ["#"]):  # Add a "#" to make program handle last part without special check.
        if line.startswith('#'):
            chunk = Chunk(
                doc_id=doc_id,
                text='\n'.join(lines[previous_lineno:lineno]).strip(),
                chunk_type=ChunkType.section,
                start_lineno=previous_lineno,
                end_lineno=lineno
            )
            current_level = line_level(lines[chunk.start_lineno])
            while len(chunk_stack) > 0 and line_level(lines[chunk_stack[-1].start_lineno]) >= current_level:
                chunk_stack.pop()

            if len(chunk_stack) > 0:
                chunk.parent_chunk_id = chunk_stack[-1].chunk_id
            if current_level > 0:
                chunk_stack.append(chunk)

            chunk_list.append(chunk)
            previous_lineno = lineno

    return chunk_list

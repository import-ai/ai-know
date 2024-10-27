import os
import pathlib
from typing import TextIO


class Resource:
    def __init__(self):
        self.resource_dir = "resource"
        self.project_root = pathlib.Path(__file__).parent
        while self.resource_dir not in os.listdir(str(self.project_root)):
            self.project_root = self.project_root.parent

    def path(self, path: str) -> str:
        return os.path.join(self.project_root, str(self.resource_dir), path)

    def open(self, path: str) -> TextIO:
        return open(self.path(path))


resource = Resource()

__all__ = [resource]

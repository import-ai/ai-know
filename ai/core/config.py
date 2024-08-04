import argparse
import os
from typing import Dict

import yaml
from pydantic import BaseModel, Field

from core.logger import get_logger

logger = get_logger(__name__)


class Config(BaseModel):
    port: int = Field(default=8000)
    workers: int = Field(default=1)

    openai_api_key: str
    openai_model: str = Field(default="gpt-3.5-turbo")
    openai_base_url: str = Field(default="https://api.openai.com/v1")

    embedding_model_name_or_path: str = Field(default="BAAI/bge-m3")
    embedding_batch_size: int = Field(default=1)
    embedding_device: str = Field(default="cpu")

    ranker_model_name_or_path: str = Field(default="BAAI/bge-reranker-v2-m3")
    ranker_batch_size: int = Field(default=1)
    ranker_device: str = Field(default="cpu")

    data_dir: str = Field(default="chroma_data")


def load_from_config_file(yaml_path: str = "config.yaml") -> Dict[str, str]:
    if os.path.exists(yaml_path):
        with open(yaml_path) as f:
            return yaml.safe_load(f)
    return {}


def load_from_env() -> Dict[str, str]:
    c: Dict[str, str] = {}

    for field_name in Config.__fields__.keys():
        c[field_name] = os.getenv(field_name.upper(), None)
    return {k: v for k, v in c.items() if v is not None}


def load_from_cli() -> Dict[str, str]:
    parser = argparse.ArgumentParser()
    for name, field in Config.__fields__.items():
        parser.add_argument(
            f"--{name}",
            dest=name,
            type=field.annotation,
            default=None,
            help=field.description,
        )
    args, _ = parser.parse_known_args()
    c = vars(args)
    return {k: v for k, v in c.items() if v is not None}


def load_config() -> Config:
    env_config: Dict[str, str] = load_from_env()
    yaml_config: Dict[str, str] = load_from_config_file()
    cli_config: Dict[str, str] = load_from_cli()
    config_merge: Dict[str, str] = {**yaml_config, **env_config, **cli_config}
    return Config.model_validate(config_merge)


config = load_config()

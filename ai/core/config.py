import argparse
import os
from typing import Dict, Type

import yaml
from pydantic import BaseModel, Field

from core.util.logger import get_logger

logger = get_logger(__name__)


class OpenAIConfig(BaseModel):
    api_key: str
    model: str = Field(default="gpt-3.5-turbo")
    base_url: str = Field(default="https://api.openai.com/v1")


class VectorDBConfig(BaseModel):
    model_name_or_path: str = Field(default="BAAI/bge-m3")
    batch_size: int = Field(default=1, description="batch size of embedding when insert data")
    device: str = Field(default="cpu")
    path: str = Field(default="chroma_data")


class RankerConfig(BaseModel):
    model_name_or_path: str = Field(default="BAAI/bge-reranker-v2-m3")
    batch_size: int = Field(default=1, description="batch size of calculating similarity")
    device: str = Field(default="cpu")


class Config(BaseModel):
    port: int = Field(default=8000)
    workers: int = Field(default=1)

    openai: OpenAIConfig
    vector_db: VectorDBConfig
    ranker: RankerConfig


def load_from_config_file(yaml_path: str = "config.yaml") -> Dict[str, str]:
    if os.path.exists(yaml_path):
        with open(yaml_path) as f:
            return yaml.safe_load(f)
    return {}


def dict_prefix_filter(prefix: str, data: dict) -> dict:
    return {k[len(prefix):]: v for k, v in data.items() if k.startswith(prefix)}


def dfs(define: Type[BaseModel], env_dict: Dict[str, str]) -> dict:
    result = {}
    for field_name, field_info in define.model_fields.items():
        filtered_env_dict = dict_prefix_filter(field_name.upper(), env_dict)
        if "" in filtered_env_dict:
            assert len(filtered_env_dict) == 1, f"Conflict name: {field_name}"
            value = filtered_env_dict.pop("")
            result[field_name] = field_info.annotation(value)
            continue
        if filtered_env_dict:
            assert issubclass(field_info.annotation, BaseModel)
            result[field_name] = dfs(field_info.annotation, dict_prefix_filter("_", filtered_env_dict))
    return result


def load_from_env() -> Dict[str, str]:
    env_prefix = "AI_KNOW"

    env_dict: Dict[str, str] = dict_prefix_filter(env_prefix, dict(os.environ))
    if "" in env_dict:
        env_dict.pop("")

    result = dfs(Config, dict_prefix_filter("_", env_dict))

    return result


def load_from_cli() -> Dict[str, str]:
    parser = argparse.ArgumentParser()
    for name, field in Config.model_fields.items():
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


__all__ = [Config, load_config, OpenAIConfig, VectorDBConfig, RankerConfig]

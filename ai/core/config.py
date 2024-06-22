import os

from pydantic import BaseModel, Field
from typing import Dict
import yaml


class Config(BaseModel):
    openai_api_key: str
    openai_model: str = Field(default="gpt-3.5-turbo")
    openai_base_url: str = Field(default="https://api.openai.com/v1")

    embedding_model_name_or_path: str = Field(default="")


def load_from_config_file(yaml_path: str = "config.yaml") -> Dict[str, str]:
    with open(yaml_path) as f:
        return yaml.safe_load(f)


def load_from_env() -> Dict[str, str]:
    config: Dict[str, str] = {}

    for field_name in Config.__fields__.keys():
        if field_name.upper() in os.environ:
            config[field_name] = os.environ[field_name]
    return config


def load_from_cli() -> Dict[str, str]:
    return {}


def load_config() -> Config:
    return Config.model_validate({**load_from_env(), **load_from_config_file(), **load_from_cli()})

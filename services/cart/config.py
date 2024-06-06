from functools import lru_cache
import os
from pydantic_settings import BaseSettings
from pydantic_settings import SettingsConfigDict


class CommonSettings(BaseSettings):
    environment: str = "dev"

    mongo_uri: str
    auth_service_port: int
    auth_service_host: str
    product_service_port: int
    cart_service_port: int

    model_config = SettingsConfigDict(env_file=".env")


class DevelopmentSettings(CommonSettings):
    class Config:
        env_file: str = ".env.development"
        extra: str = "allow"


class ProductionSettings(CommonSettings):
    class Config:
        env_file: str = ".env.production"
        extra: str = "allow"


@lru_cache()
def get_settings() -> CommonSettings:
    """
    Returns cached settings object based on the current environment.
    """
    environment = os.getenv("ENVIRONMENT", "dev")

    if environment == "prod":
        return ProductionSettings()

    return DevelopmentSettings()


settings = get_settings()

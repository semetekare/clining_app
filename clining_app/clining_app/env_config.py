from pydantic_settings import BaseSettings


class Config(BaseSettings):
    DB_NAME: str = 'postgres'
    DB_USER: str = 'postgres'
    DB_PASSWORD: str = '3803'
    DB_HOST: str = 'localhost'
    DB_PORT: int = 5432

CONFIG = Config()

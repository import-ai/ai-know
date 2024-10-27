import uvicorn

from core.config import Config, load_config

config: Config = load_config()


def main():
    uvicorn.run(
        'app:app',
        host='0.0.0.0',
        port=config.port,
        workers=config.workers
    )


if __name__ == '__main__':
    main()
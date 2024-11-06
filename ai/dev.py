from app import app
import uvicorn

from core.config import Config, load_config

config: Config = load_config()

@app.get("/dev/")
async def index():
    pass

def main():
    uvicorn.run(
        'dev:app',
        host='0.0.0.0',
        port=config.port,
        workers=config.workers,
        reload=True
    )


if __name__ == '__main__':
    main()

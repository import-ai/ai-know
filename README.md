# AIKnow

## Deploy

1. Create `ai/.env` file using the following template with `OPENAI_BASE_URL` and
`OPENAI_API_KEY` filled in:
```sh
OPENAI_BASE_URL=
OPENAI_API_KEY=
OPENAI_MODEL=gpt-4o
DEVICE=cpu
WORKERS=1
DATA_DIR=chroma_data
```
2. `docker compose up --build`
3. http://localhost:8911

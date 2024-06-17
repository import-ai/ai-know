SET hnsw_enable_experimental_persistence = true;

CREATE TYPE VECTOR AS FLOAT[3];

CREATE TABLE IF NOT EXISTS vectors (
    doc_id INTEGER,
    vec VECTOR
);

CREATE INDEX IF NOT EXISTS cos_idx
    ON vectors
    USING HNSW (vec)
    WITH (metric = 'cosine')
;


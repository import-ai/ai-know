CREATE TYPE sidebar_entry_type AS ENUM ('group','note', 'link');

CREATE TABLE sidebar_entries
(
    id         BIGSERIAL                NOT NULL PRIMARY KEY,
    type       sidebar_entry_type       NOT NULL,
    title      TEXT                     NOT NULL,
    parent_id  BIGINT                   NULL REFERENCES sidebar_entries (id),
    prev_id    BIGINT                   NULL REFERENCES sidebar_entries (id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX uniq_sidebar_par_prev ON sidebar_entries (parent_id, prev_id);

CREATE TABLE workspaces
(
    id                    BIGSERIAL                NOT NULL PRIMARY KEY,
    private_sidebar_entry BIGINT                   NOT NULL REFERENCES sidebar_entries (id),
    team_sidebar_entry    BIGINT                   NOT NULL REFERENCES sidebar_entries (id),
    created_at            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
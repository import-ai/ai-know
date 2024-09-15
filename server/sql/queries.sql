-- name: CreateWorkspace :one
INSERT INTO workspaces(id, private_sidebar_entry, team_sidebar_entry)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetWorkspace :one
SELECT *
FROM workspaces
WHERE id = $1;

-- name: CreateSidebarEntry :one
INSERT INTO sidebar_entries(type, title, parent_id, prev_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ReplacePrevEntry :one
UPDATE sidebar_entries
SET prev_id = @new_prev
WHERE parent_id = @parent_id
  AND prev_id = @old_prev
RETURNING *;

-- name: SetPrevEntry :one
UPDATE sidebar_entries
SET prev_id = $2
WHERE id = $1
  AND prev_id IS NULL
RETURNING *;

-- name: GetSidebarEntry :one
SELECT *
FROM sidebar_entries
WHERE id = $1;

-- name: LockSidebarEntry :one
SELECT *
FROM sidebar_entries
WHERE id = $1 FOR UPDATE;

-- name: GetNextSidebarEntry :one
SELECT *
FROM sidebar_entries
WHERE parent_id = $1
  AND prev_id = $2;
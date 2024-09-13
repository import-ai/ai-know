-- name: CreateWorkspace :one
INSERT INTO workspaces(id, private_sidebar_entry, team_sidebar_entry)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetWorkspace :one
SELECT *
FROM workspaces
WHERE id = $1;

-- name: CreateSidebarEntry :one
INSERT INTO sidebar_entries(type, title, parent_id, first_child_id, next_brother_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateFirstChild :one
UPDATE sidebar_entries
SET first_child_id = @new_val
WHERE id = @id
  AND first_child_id = @old_val
RETURNING *;

-- name: UpdateNextBrother :one
UPDATE sidebar_entries
SET next_brother_id = @new_val
WHERE id = @id
  AND next_brother_id = @old_val
RETURNING *;

-- name: GetSidebarEntry :one
SELECT *
FROM sidebar_entries
WHERE id = $1;
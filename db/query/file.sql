-- name: CreateFile :one
INSERT INTO files (
  file_name,
  owner,
  chunk_count
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetFile :one
SELECT * FROM files
WHERE hash = $1 LIMIT 1;
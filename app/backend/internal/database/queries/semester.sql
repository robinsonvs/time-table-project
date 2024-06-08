-- name: GetSemesterByID :one
SELECT * from semester s where s.uuid = $1;

-- name: CreateSemester :exec
INSERT INTO semester (uuid, semester)
VALUES ($1, $2);

-- name: FindSemesterByID :one
SELECT s.id, s.uuid, s.semester
FROM semester s
WHERE s.uuid = $1;

-- name: UpdateSemester :exec
UPDATE semester SET
    semester = COALESCE(sqlc.narg('semester'), semester)
WHERE uuid = $1;

-- name: DeleteSemester :exec
DELETE FROM semester WHERE uuid = $1;

-- name: FindManySemesters :many
SELECT s.id, s.uuid, s.semester
FROM semester s
ORDER BY s.semester ASC;



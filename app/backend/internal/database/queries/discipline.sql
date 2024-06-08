-- name: GetDisciplineByID :one
SELECT * from discipline d where d.uuid = $1;

-- name: CreateDiscipline :exec
INSERT INTO discipline (uuid, name, credits, course_id)
VALUES ($1, $2, $3, $4);

-- name: FindDisciplineByID :one
SELECT d.id, d.uuid, d.name, d.credits, d.course_id
FROM discipline d
WHERE d.uuid = $1;

-- name: UpdateDiscipline :exec
UPDATE discipline SET
    name = COALESCE(sqlc.narg('name'), name),
    credits = COALESCE(sqlc.narg('credits'), credits)
WHERE uuid = $1;

-- name: DeleteDiscipline :exec
DELETE FROM discipline WHERE uuid = $1;

-- name: FindManyDisciplines :many
SELECT d.id, d.uuid, d.name, d.credits, d.course_id
FROM discipline d
ORDER BY d.course_id, d.name ASC;

-- name: FindManyDisciplinesByCourseId :many
SELECT d.id, d.uuid, d.name, d.credits, d.course_id
FROM discipline d
WHERE d.course_id = $1
ORDER BY d.name ASC;



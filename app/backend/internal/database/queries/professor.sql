-- name: GetProfessorByID :one
SELECT * from professor p where p.uuid = $1;

-- name: CreateProfessor :exec
INSERT INTO professor (uuid, name, hoursToAllocate)
VALUES ($1, $2, $3);

-- name: FindProfessorByID :one
SELECT p.id, p.uuid, p.name, p.hoursToAllocate
FROM professor p
WHERE p.uuid = $1;

-- name: UpdateProfessor :exec
UPDATE professor SET
    name = COALESCE(sqlc.narg('name'), semester),
    hoursToAllocate = COALESCE(sqlc.narg('hoursToAllocate'), hoursToAllocate)
WHERE uuid = $1;

-- name: DeleteProfessor :exec
DELETE FROM professor WHERE uuid = $1;

-- name: FindManyProfessors :many
SELECT p.id, p.uuid, p.name, p.hoursToAllocate
FROM professor p
ORDER BY p.name ASC;



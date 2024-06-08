-- name: GetAvailabilityByID :one
SELECT * from availability a where a.uuid = $1;

-- name: CreateAvailability :exec
INSERT INTO availability (uuid, dayOfWeek, shift, professor_id)
VALUES ($1, $2, $3, $4);

-- name: FindAvailabilityByID :one
SELECT a.id, a.uuid, a.dayOfWeek, a.shift, a.professor_id
FROM availability a
WHERE a.uuid = $1;

-- name: UpdateAvailability :exec
UPDATE availability SET
    dayOfWeek = COALESCE(sqlc.narg('dayOfWeek'), dayOfWeek),
    shift = COALESCE(sqlc.narg('shift'), shift)
WHERE uuid = $1;

-- name: DeleteAvailability :exec
DELETE FROM availability WHERE uuid = $1;

-- name: FindManyAvailabilities :many
SELECT a.id, a.uuid, a.dayOfWeek, a.shift, a.professor_id
FROM availability a
ORDER BY a.dayOfWeek, a.shift ASC;

-- name: FindManyAvailabilitiesByProfessorId :many
SELECT a.id, a.uuid, a.dayOfWeek, a.shift, a.professor_id
FROM availability a
WHERE a.professor_id = $1
ORDER BY a.professor_id, a.dayOfWeek, a.shift ASC;



-- name: GetParameterizationByID :one
SELECT * from parameterization p where p.uuid = $1;

-- name: CreateParameterization :exec
INSERT INTO parameterization (uuid, maxCreditsToOffer, numClassesPerDiscipline, semester_id, course_id)
VALUES ($1, $2, $3, $4, $5);

-- name: FindParameterizationByID :one
SELECT p.id, p.uuid, p.maxCreditsToOffer, p.numClassesPerDiscipline, p.semester_id, p.course_id
FROM parameterization p
WHERE p.uuid = $1;

-- name: UpdateParameterization :exec
UPDATE parameterization SET
    maxCreditsToOffer = COALESCE(sqlc.narg('maxCreditsToOffer'), maxCreditsToOffer),
    numClassesPerDiscipline = COALESCE(sqlc.narg('numClassesPerDiscipline'), numClassesPerDiscipline)
WHERE uuid = $1;

-- name: DeleteParameterization :exec
DELETE FROM parameterization WHERE uuid = $1;

-- name: FindManyParameterizations :many
SELECT p.id, p.uuid, p.maxCreditsToOffer, p.numClassesPerDiscipline, p.semester_id, p.course_id
FROM parameterization p
ORDER BY p.semester_id, p.course_id ASC;

-- name: FindManyParameterizationsBySemesterId :many
SELECT p.id, p.uuid, p.maxCreditsToOffer, p.numClassesPerDiscipline, p.semester_id, p.course_id
FROM parameterization p
WHERE p.semester_id = $1
ORDER BY p.course_id ASC;



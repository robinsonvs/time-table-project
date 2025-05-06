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




-- name: GetDisciplinesByCourseID :many
SELECT id, uuid, name, credits, course_id FROM discipline WHERE course_id = $1;

-- name: GetProfessorsByCourseID :many
SELECT p.id, p.uuid, p.name, p.hoursToAllocate
FROM professor p
         JOIN eligible_disciplines ed ON p.id = ed.professor_id
         JOIN discipline d ON ed.discipline_id = d.id
WHERE d.course_id = $1;

-- name: CreateProposal :exec
INSERT INTO proposal (uuid, semester_id, course_id)
VALUES ($1, $2, $3);

-- name: CreateClass :exec
INSERT INTO class (uuid, dayOfWeek, shift, startTime, endTime, discipline_id, professor_id, proposal_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, uuid;

-- name: GetProposalID :one
SELECT p.id from proposal p where p.uuid = $1;


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
    name = COALESCE(sqlc.narg('name'), name),
    hoursToAllocate = COALESCE(sqlc.narg('hoursToAllocate'), hoursToAllocate)
WHERE uuid = $1;

-- name: DeleteProfessor :exec
DELETE FROM professor WHERE uuid = $1;

-- name: FindManyProfessors :many
SELECT p.id, p.uuid, p.name, p.hoursToAllocate
FROM professor p
ORDER BY p.name ASC;

-- name: GetProfessorsWithDisciplines :many
SELECT
    p.id AS professor_id,
    p.uuid AS professor_uuid,
    p.name AS professor_name,
    p.hoursToAllocate AS professor_hours_to_allocate,
    d.id AS discipline_id,
    d.uuid AS discipline_uuid,
    d.name AS discipline_name,
    d.credits AS discipline_credits,
    d.course_id AS discipline_course_id
FROM
    professor p
        LEFT JOIN
    eligible_disciplines ed ON p.id = ed.professor_id
        LEFT JOIN
    discipline d ON ed.discipline_id = d.id;



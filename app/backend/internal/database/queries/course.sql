-- name: GetCourseByID :one
SELECT * from course c where c.uuid = $1;

-- name: CreateCourse :exec
INSERT INTO course (uuid, name, modality, location)
VALUES ($1, $2, $3, $4);

-- name: FindCourseByID :one
SELECT c.id, c.uuid, c.name, c.modality, c.location
FROM course c
WHERE c.uuid = $1;

-- name: UpdateCourse :exec
UPDATE course SET
                 name = COALESCE(sqlc.narg('name'), name),
                 modality = COALESCE(sqlc.narg('modality'), modality),
                 location = COALESCE(sqlc.narg('location'), location)
WHERE uuid = $1;

-- name: DeleteCourse :exec
DELETE FROM course WHERE uuid = $1;

-- name: FindManyCourses :many
SELECT c.id, c.uuid, c.name, c.modality, c.location
FROM course c
ORDER BY c.name ASC;



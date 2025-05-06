
-- name: CreateEligibleDiscipline :exec
INSERT INTO eligible_disciplines (professor_id, discipline_id)
VALUES ($1, $2);

-- name: DeleteEligibleDiscipline :exec
DELETE FROM eligible_disciplines WHERE professor_id = $1 AND discipline_id = $2;



// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: availability.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createAvailability = `-- name: CreateAvailability :exec
INSERT INTO availability (uuid, dayOfWeek, shift, professor_id)
VALUES ($1, $2, $3, $4)
`

type CreateAvailabilityParams struct {
	Uuid        uuid.UUID
	Dayofweek   string
	Shift       string
	ProfessorID int64
}

func (q *Queries) CreateAvailability(ctx context.Context, arg CreateAvailabilityParams) error {
	_, err := q.db.ExecContext(ctx, createAvailability,
		arg.Uuid,
		arg.Dayofweek,
		arg.Shift,
		arg.ProfessorID,
	)
	return err
}

const deleteAvailability = `-- name: DeleteAvailability :exec
DELETE FROM availability WHERE uuid = $1
`

func (q *Queries) DeleteAvailability(ctx context.Context, argUuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAvailability, argUuid)
	return err
}

const findAvailabilityByID = `-- name: FindAvailabilityByID :one
SELECT a.id, a.uuid, a.dayOfWeek, a.shift, a.professor_id
FROM availability a
WHERE a.uuid = $1
`

func (q *Queries) FindAvailabilityByID(ctx context.Context, argUuid uuid.UUID) (Availability, error) {
	row := q.db.QueryRowContext(ctx, findAvailabilityByID, argUuid)
	var i Availability
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Dayofweek,
		&i.Shift,
		&i.ProfessorID,
	)
	return i, err
}

const findManyAvailabilities = `-- name: FindManyAvailabilities :many
SELECT a.id, a.uuid, a.dayOfWeek, a.shift, a.professor_id
FROM availability a
ORDER BY a.dayOfWeek, a.shift ASC
`

func (q *Queries) FindManyAvailabilities(ctx context.Context) ([]Availability, error) {
	rows, err := q.db.QueryContext(ctx, findManyAvailabilities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Availability
	for rows.Next() {
		var i Availability
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.Dayofweek,
			&i.Shift,
			&i.ProfessorID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findManyAvailabilitiesByProfessorId = `-- name: FindManyAvailabilitiesByProfessorId :many
SELECT a.id, a.uuid, a.dayOfWeek, a.shift, a.professor_id
FROM availability a
WHERE a.professor_id = $1
ORDER BY a.professor_id, a.dayOfWeek, a.shift ASC
`

func (q *Queries) FindManyAvailabilitiesByProfessorId(ctx context.Context, professorID int64) ([]Availability, error) {
	rows, err := q.db.QueryContext(ctx, findManyAvailabilitiesByProfessorId, professorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Availability
	for rows.Next() {
		var i Availability
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.Dayofweek,
			&i.Shift,
			&i.ProfessorID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAvailabilityByID = `-- name: GetAvailabilityByID :one
SELECT id, uuid, dayofweek, shift, professor_id from availability a where a.uuid = $1
`

func (q *Queries) GetAvailabilityByID(ctx context.Context, argUuid uuid.UUID) (Availability, error) {
	row := q.db.QueryRowContext(ctx, getAvailabilityByID, argUuid)
	var i Availability
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Dayofweek,
		&i.Shift,
		&i.ProfessorID,
	)
	return i, err
}

const updateAvailability = `-- name: UpdateAvailability :exec
UPDATE availability SET
    dayOfWeek = COALESCE($2, dayOfWeek),
    shift = COALESCE($3, shift)
WHERE uuid = $1
`

type UpdateAvailabilityParams struct {
	Uuid      uuid.UUID
	DayOfWeek sql.NullString
	Shift     sql.NullString
}

func (q *Queries) UpdateAvailability(ctx context.Context, arg UpdateAvailabilityParams) error {
	_, err := q.db.ExecContext(ctx, updateAvailability, arg.Uuid, arg.DayOfWeek, arg.Shift)
	return err
}

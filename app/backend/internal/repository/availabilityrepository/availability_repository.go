package availabilityrepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func (r *repository) CreateAvailability(ctx context.Context, u *entity.AvailabilityEntity) error {
	err := r.queries.CreateAvailability(ctx, sqlc.CreateAvailabilityParams{
		Uuid:        u.UUID,
		Dayofweek:   u.DayOfWeek,
		Shift:       u.Shift,
		ProfessorID: u.ProfessorID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAvailabilityByID(ctx context.Context, uuid uuid.UUID) (*entity.AvailabilityEntity, error) {
	availability, err := r.queries.FindAvailabilityByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	availabilityEntity := entity.AvailabilityEntity{
		UUID:        availability.Uuid,
		DayOfWeek:   availability.Dayofweek,
		Shift:       availability.Shift,
		ProfessorID: availability.ProfessorID,
	}

	return &availabilityEntity, nil
}

func (r *repository) UpdateAvailability(ctx context.Context, u *entity.AvailabilityEntity) error {
	err := r.queries.UpdateAvailability(ctx, sqlc.UpdateAvailabilityParams{
		Uuid:      u.UUID,
		DayOfWeek: sql.NullString{String: u.DayOfWeek, Valid: u.DayOfWeek != ""},
		Shift:     sql.NullString{String: u.Shift, Valid: u.Shift != ""},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteAvailability(ctx context.Context, uuid uuid.UUID) error {
	err := r.queries.DeleteAvailability(ctx, uuid)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindManyAvailabilities(ctx context.Context) ([]entity.AvailabilityEntity, error) {
	availabilities, err := r.queries.FindManyAvailabilities(ctx)
	if err != nil {
		return nil, err
	}

	var availabilitiesEntity []entity.AvailabilityEntity
	for _, availability := range availabilities {
		availabilityEntity := entity.AvailabilityEntity{
			UUID:        availability.Uuid,
			DayOfWeek:   availability.Dayofweek,
			Shift:       availability.Shift,
			ProfessorID: availability.ProfessorID,
		}

		availabilitiesEntity = append(availabilitiesEntity, availabilityEntity)
	}
	return availabilitiesEntity, nil
}

func (r *repository) FindManyAvailabilitiesByProfessorId(ctx context.Context, professorId int64) ([]entity.AvailabilityEntity, error) {
	availabilities, err := r.queries.FindManyAvailabilitiesByProfessorId(ctx, professorId)
	if err != nil {
		return nil, err
	}

	var availabilitiesEntity []entity.AvailabilityEntity
	for _, availability := range availabilities {
		availabilityEntity := entity.AvailabilityEntity{
			UUID:        availability.Uuid,
			DayOfWeek:   availability.Dayofweek,
			Shift:       availability.Shift,
			ProfessorID: availability.ProfessorID,
		}

		availabilitiesEntity = append(availabilitiesEntity, availabilityEntity)
	}
	return availabilitiesEntity, nil
}

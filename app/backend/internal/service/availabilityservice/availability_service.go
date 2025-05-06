package availabilityservice

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/entity"
	"github.com/robinsonvs/time-table-project/internal/handler/response"
	"log/slog"
)

func (s *service) CreateAvailability(ctx context.Context, u dto.CreateAvailabilityDto) error {

	newAvailability := entity.AvailabilityEntity{
		UUID:        uuid.New(),
		DayOfWeek:   u.DayOfWeek,
		Shift:       u.Shift,
		ProfessorID: u.ProfessorId,
	}

	err := s.repo.CreateAvailability(ctx, &newAvailability)
	if err != nil {
		slog.Error("error to create availability", "err", err, slog.String("package", "availabilityservice"))
		return err
	}

	return nil
}

func (s *service) UpdateAvailability(ctx context.Context, u dto.UpdateAvailabilityDto, uuid uuid.UUID) error {
	availabilityExists, err := s.repo.FindAvailabilityByID(ctx, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("availability not found", slog.String("package", "availabilityservice"))
			return errors.New("availability not found")
		}
		slog.Error("error to search availability by id", "err", err, slog.String("package", "availabilityservice"))
		return err
	}

	if availabilityExists == nil {
		slog.Error("availability not found", slog.String("package", "availabilityservice"))
		return errors.New("availability already exists")
	}

	updateAvailability := entity.AvailabilityEntity{
		UUID:      uuid,
		DayOfWeek: u.DayOfWeek,
		Shift:     u.Shift,
	}

	err = s.repo.UpdateAvailability(ctx, &updateAvailability)
	if err != nil {
		slog.Error("error to update discipline", "err", err, slog.String("package", "disciplineservice"))
		return err
	}

	return nil
}

func (s *service) GetAvailabilityByID(ctx context.Context, uuid uuid.UUID) (*response.AvailabilityResponse, error) {
	availabilityExists, err := s.repo.FindAvailabilityByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search availability by id", "err", err, slog.String("package", "availabilityservice"))
		return nil, err
	}

	if availabilityExists == nil {
		slog.Error("availability not found", slog.String("package", "availabilityservice"))
		return nil, errors.New("availability not found")
	}

	availability := response.AvailabilityResponse{
		UUID:        availabilityExists.UUID.String(),
		DayOfWeek:   availabilityExists.DayOfWeek,
		Shift:       availabilityExists.Shift,
		ProfessorId: availabilityExists.ProfessorID,
	}

	return &availability, nil
}

func (s *service) FindManyAvailabilities(ctx context.Context) (*response.ManyAvailabilitiesResponse, error) {
	findManyAvailabilities, err := s.repo.FindManyAvailabilities(ctx)
	if err != nil {
		slog.Error("error to find many availabilities", "err", err, slog.String("package", "availabilitiesservice"))
		return nil, err
	}

	availabilities := response.ManyAvailabilitiesResponse{}
	for _, availabilityEntity := range findManyAvailabilities {
		availabilityResponse := response.AvailabilityResponse{
			Id:          availabilityEntity.ID,
			UUID:        availabilityEntity.UUID.String(),
			DayOfWeek:   availabilityEntity.DayOfWeek,
			Shift:       availabilityEntity.Shift,
			ProfessorId: availabilityEntity.ProfessorID,
		}
		availabilities.Availabilities = append(availabilities.Availabilities, availabilityResponse)
	}

	return &availabilities, nil
}

func (s *service) FindManyAvailabilitiesByProfessorId(ctx context.Context, courseId int64) (*response.ManyAvailabilitiesResponse, error) {
	findManyAvailabilitiesByProfessor, err := s.repo.FindManyAvailabilitiesByProfessorId(ctx, courseId)
	if err != nil {
		slog.Error("error to find many availabilities", "err", err, slog.String("package", "availabilityservice"))
		return nil, err
	}

	availabilitiesByProfessor := response.ManyAvailabilitiesResponse{}
	for _, availabilityEntity := range findManyAvailabilitiesByProfessor {
		availabilityResponse := response.AvailabilityResponse{
			Id:          availabilityEntity.ID,
			UUID:        availabilityEntity.UUID.String(),
			DayOfWeek:   availabilityEntity.DayOfWeek,
			Shift:       availabilityEntity.Shift,
			ProfessorId: availabilityEntity.ProfessorID,
		}
		availabilitiesByProfessor.Availabilities = append(availabilitiesByProfessor.Availabilities, availabilityResponse)
	}

	return &availabilitiesByProfessor, nil
}

func (s *service) DeleteAvailability(ctx context.Context, uuid uuid.UUID) error {
	availabilityExists, err := s.repo.FindAvailabilityByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search availability id", "err", err, slog.String("package", "availabilityservice"))
		return err
	}

	if availabilityExists == nil {
		slog.Error("availability not found", slog.String("package", "availabilityservice"))
		return errors.New("availability not found")
	}

	err = s.repo.DeleteAvailability(ctx, uuid)
	if err != nil {
		slog.Error("error to delete availability", "err", err, slog.String("package", "availabilityservice"))
		return err
	}

	return nil
}

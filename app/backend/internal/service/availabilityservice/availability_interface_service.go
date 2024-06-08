package availabilityservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/handler/response"
	"github.com/robinsonvs/time-table-project/internal/repository/availabilityrepository"
)

func NewAvailabilityService(repo availabilityrepository.AvailabilityRepository) AvailabilityService {
	return &service{
		repo,
	}
}

type service struct {
	repo availabilityrepository.AvailabilityRepository
}

type AvailabilityService interface {
	CreateAvailability(ctx context.Context, u dto.CreateAvailabilityDto) error
	UpdateAvailability(ctx context.Context, u dto.UpdateAvailabilityDto, uuid uuid.UUID) error
	GetAvailabilityByID(ctx context.Context, uuid uuid.UUID) (*response.AvailabilityResponse, error)
	DeleteAvailability(ctx context.Context, uuid uuid.UUID) error
	FindManyAvailabilities(ctx context.Context) (*response.ManyAvailabilitiesResponse, error)
	FindManyAvailabilitiesByProfessorId(ctx context.Context, professorId int64) (*response.ManyAvailabilitiesResponse, error)
}

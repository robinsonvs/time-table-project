package parameterizationservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/handler/response"
	"github.com/robinsonvs/time-table-project/internal/repository/parameterizationrepository"
)

func NewParameterizationService(repo parameterizationrepository.ParameterizationRepository) ParameterizationService {
	return &service{
		repo,
	}
}

type service struct {
	repo parameterizationrepository.ParameterizationRepository
}

type ParameterizationService interface {
	CreateParameterization(ctx context.Context, u dto.CreateParameterizationDto) error
	UpdateParameterization(ctx context.Context, u dto.UpdateParameterizationDto, uuid uuid.UUID) error
	GetParameterizationByID(ctx context.Context, uuid uuid.UUID) (*response.ParameterizationResponse, error)
	DeleteParameterization(ctx context.Context, uuid uuid.UUID) error
	FindManyParameterizations(ctx context.Context) (*response.ManyParameterizationsResponse, error)
	FindManyParameterizationsBySemesterId(ctx context.Context, semesterId int64) (*response.ManyParameterizationsResponse, error)
}

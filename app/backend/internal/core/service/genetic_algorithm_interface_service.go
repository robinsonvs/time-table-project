package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/repository/availabilityrepository"
	"github.com/robinsonvs/time-table-project/internal/repository/disciplinerepository"
	"github.com/robinsonvs/time-table-project/internal/repository/parameterizationrepository"
	"github.com/robinsonvs/time-table-project/internal/repository/professorrepository"
)

func NewGeneticAlgorithmService(
	disciplineRepo disciplinerepository.DisciplineRepository,
	professorRepo professorrepository.ProfessorRepository,
	availabilityRepo availabilityrepository.AvailabilityRepository,
	parameterizationRepo parameterizationrepository.ParameterizationRepository,
) GeneticAlgorithmServiceInterface {
	return &GeneticAlgorithmService{
		DisciplineRepo:       disciplineRepo,
		ProfessorRepo:        professorRepo,
		AvailabilityRepo:     availabilityRepo,
		ParameterizationRepo: parameterizationRepo,
	}
}

type GeneticAlgorithmService struct {
	DisciplineRepo       disciplinerepository.DisciplineRepository
	ProfessorRepo        professorrepository.ProfessorRepository
	AvailabilityRepo     availabilityrepository.AvailabilityRepository
	ParameterizationRepo parameterizationrepository.ParameterizationRepository
}

type GeneticAlgorithmServiceInterface interface {
	GenerateProposal(ctx context.Context, parameterizationID uuid.UUID) error
}

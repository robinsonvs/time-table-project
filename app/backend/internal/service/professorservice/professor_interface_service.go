package professorservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/handler/response"
	"github.com/robinsonvs/time-table-project/internal/repository/professorrepository"
)

func NewProfessorService(repo professorrepository.ProfessorRepository) ProfessorService {
	return &service{
		repo,
	}
}

type service struct {
	repo professorrepository.ProfessorRepository
}

type ProfessorService interface {
	CreateProfessor(ctx context.Context, u dto.CreateProfessorDto) error
	UpdateProfessor(ctx context.Context, u dto.UpdateProfessorDto, uuid uuid.UUID) error
	GetProfessorByID(ctx context.Context, uuid uuid.UUID) (*response.ProfessorResponse, error)
	DeleteProfessor(ctx context.Context, uuid uuid.UUID) error
	FindManyProfessors(ctx context.Context) (*response.ManyProfessorsResponse, error)
}

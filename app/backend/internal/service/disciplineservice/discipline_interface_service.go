package disciplineservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/handler/response"
	"github.com/robinsonvs/time-table-project/internal/repository/disciplinerepository"
)

func NewDisciplineService(repo disciplinerepository.DisciplineRepository) DisciplineService {
	return &service{
		repo,
	}
}

type service struct {
	repo disciplinerepository.DisciplineRepository
}

type DisciplineService interface {
	CreateDiscipline(ctx context.Context, u dto.CreateDisciplineDto) error
	UpdateDiscipline(ctx context.Context, u dto.UpdateDisciplineDto, uuid uuid.UUID) error
	GetDisciplineByID(ctx context.Context, uuid uuid.UUID) (*response.DisciplineResponse, error)
	DeleteDiscipline(ctx context.Context, uuid uuid.UUID) error
	FindManyDisciplines(ctx context.Context) (*response.ManyDisciplinesResponse, error)
	FindManyDisciplinesByCourseId(ctx context.Context, courseId int64) (*response.ManyDisciplinesResponse, error)
}

package semesterservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/handler/response"
	"github.com/robinsonvs/time-table-project/internal/repository/semesterrepository"
)

func NewSemesterService(repo semesterrepository.SemesterRepository) SemesterService {
	return &service{
		repo,
	}
}

type service struct {
	repo semesterrepository.SemesterRepository
}

type SemesterService interface {
	CreateSemester(ctx context.Context, u dto.CreateSemesterDto) error
	UpdateSemester(ctx context.Context, u dto.UpdateSemesterDto, uuid uuid.UUID) error
	GetSemesterByID(ctx context.Context, uuid uuid.UUID) (*response.SemesterResponse, error)
	DeleteSemester(ctx context.Context, uuid uuid.UUID) error
	FindManySemesters(ctx context.Context) (*response.ManySemestersResponse, error)
}

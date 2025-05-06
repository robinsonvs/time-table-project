package eligibledisciplineservice

import (
	"context"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/repository/eligibledisciplinerepository"
)

func NewEligibleDisciplineService(repo eligibledisciplinerepository.EligibleDisciplineRepository) EligibleDisciplineService {
	return &service{
		repo,
	}
}

type service struct {
	repo eligibledisciplinerepository.EligibleDisciplineRepository
}

type EligibleDisciplineService interface {
	CreateEligibleDiscipline(ctx context.Context, u dto.CreateEligibleDisciplineDto) error
	DeleteEligibleDiscipline(ctx context.Context, u dto.DeleteEligibleDisciplineDto) error
}

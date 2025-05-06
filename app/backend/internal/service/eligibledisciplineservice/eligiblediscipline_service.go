package eligibledisciplineservice

import (
	"context"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/entity"
	"log/slog"
)

func (s *service) CreateEligibleDiscipline(ctx context.Context, u dto.CreateEligibleDisciplineDto) error {

	newEligibleDiscipline := entity.EligibleDisciplineEntity{
		ProfessorID:  u.ProfessorId,
		DisciplineID: u.DisciplineId,
	}

	err := s.repo.CreateEligibleDiscipline(ctx, &newEligibleDiscipline)
	if err != nil {
		slog.Error("error to create eligible discipline", "err", err, slog.String("package", "eligibledisciplineservice"))
		return err
	}

	return nil
}

func (s *service) DeleteEligibleDiscipline(ctx context.Context, u dto.DeleteEligibleDisciplineDto) error {

	deleteEligibleDiscipline := entity.EligibleDisciplineEntity{
		ProfessorID:  u.ProfessorId,
		DisciplineID: u.DisciplineId,
	}

	err := s.repo.DeleteEligibleDiscipline(ctx, &deleteEligibleDiscipline)
	if err != nil {
		slog.Error("error to delete eligible discipline", "err", err, slog.String("package", "eligibledisciplineservice"))
		return err
	}

	return nil
}

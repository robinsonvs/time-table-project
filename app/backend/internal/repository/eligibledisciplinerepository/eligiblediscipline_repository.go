package eligibledisciplinerepository

import (
	"context"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func (r *repository) CreateEligibleDiscipline(ctx context.Context, u *entity.EligibleDisciplineEntity) error {
	err := r.queries.CreateEligibleDiscipline(ctx, sqlc.CreateEligibleDisciplineParams{
		ProfessorID:  u.ProfessorID,
		DisciplineID: u.DisciplineID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteEligibleDiscipline(ctx context.Context, u *entity.EligibleDisciplineEntity) error {
	err := r.queries.DeleteEligibleDiscipline(ctx, sqlc.DeleteEligibleDisciplineParams{
		ProfessorID:  u.ProfessorID,
		DisciplineID: u.DisciplineID,
	})

	if err != nil {
		return err
	}

	return nil
}

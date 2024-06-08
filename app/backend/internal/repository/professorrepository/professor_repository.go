package professorrepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func (r *repository) CreateProfessor(ctx context.Context, u *entity.ProfessorEntity) error {
	err := r.queries.CreateProfessor(ctx, sqlc.CreateProfessorParams{
		Uuid:            u.UUID,
		Name:            u.Name,
		Hourstoallocate: u.HoursToAllocate,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindProfessorByID(ctx context.Context, uuid uuid.UUID) (*entity.ProfessorEntity, error) {
	professor, err := r.queries.FindProfessorByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	professorEntity := entity.ProfessorEntity{
		UUID:            professor.Uuid,
		Name:            professor.Name,
		HoursToAllocate: professor.Hourstoallocate,
	}

	return &professorEntity, nil
}

func (r *repository) UpdateProfessor(ctx context.Context, u *entity.ProfessorEntity) error {
	err := r.queries.UpdateProfessor(ctx, sqlc.UpdateProfessorParams{
		Uuid:            u.UUID,
		Name:            sql.NullString{String: u.Name, Valid: u.Name != ""},
		HoursToAllocate: sql.NullInt32{Int32: u.HoursToAllocate, Valid: u.HoursToAllocate != 0},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteProfessor(ctx context.Context, uuid uuid.UUID) error {
	err := r.queries.DeleteProfessor(ctx, uuid)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindManyProfessors(ctx context.Context) ([]entity.ProfessorEntity, error) {
	professors, err := r.queries.FindManyProfessors(ctx)
	if err != nil {
		return nil, err
	}

	var professorsEntity []entity.ProfessorEntity
	for _, professor := range professors {
		professorEntity := entity.ProfessorEntity{
			UUID:            professor.Uuid,
			Name:            professor.Name,
			HoursToAllocate: professor.Hourstoallocate,
		}

		professorsEntity = append(professorsEntity, professorEntity)
	}
	return professorsEntity, nil
}

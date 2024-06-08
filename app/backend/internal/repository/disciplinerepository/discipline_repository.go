package disciplinerepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func (r *repository) CreateDiscipline(ctx context.Context, u *entity.DisciplineEntity) error {
	err := r.queries.CreateDiscipline(ctx, sqlc.CreateDisciplineParams{
		Uuid:     u.UUID,
		Name:     u.Name,
		Credits:  u.Credits,
		CourseID: u.CourseID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindDisciplineByID(ctx context.Context, uuid uuid.UUID) (*entity.DisciplineEntity, error) {
	discipline, err := r.queries.FindDisciplineByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	disciplineEntity := entity.DisciplineEntity{
		UUID:     discipline.Uuid,
		Name:     discipline.Name,
		Credits:  discipline.Credits,
		CourseID: discipline.CourseID,
	}

	return &disciplineEntity, nil
}

func (r *repository) UpdateDiscipline(ctx context.Context, u *entity.DisciplineEntity) error {
	err := r.queries.UpdateDiscipline(ctx, sqlc.UpdateDisciplineParams{
		Uuid:    u.UUID,
		Name:    sql.NullString{String: u.Name, Valid: u.Name != ""},
		Credits: sql.NullInt32{Int32: u.Credits, Valid: u.Credits != 0},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteDiscipline(ctx context.Context, uuid uuid.UUID) error {
	err := r.queries.DeleteDiscipline(ctx, uuid)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindManyDisciplines(ctx context.Context) ([]entity.DisciplineEntity, error) {
	disciplines, err := r.queries.FindManyDisciplines(ctx)
	if err != nil {
		return nil, err
	}

	var disciplinesEntity []entity.DisciplineEntity
	for _, discipline := range disciplines {
		disciplineEntity := entity.DisciplineEntity{
			UUID:     discipline.Uuid,
			Name:     discipline.Name,
			Credits:  discipline.Credits,
			CourseID: discipline.CourseID,
		}

		disciplinesEntity = append(disciplinesEntity, disciplineEntity)
	}
	return disciplinesEntity, nil
}

func (r *repository) FindManyDisciplinesByCoarseId(ctx context.Context, courseId int64) ([]entity.DisciplineEntity, error) {
	disciplines, err := r.queries.FindManyDisciplinesByCourseId(ctx, courseId)
	if err != nil {
		return nil, err
	}

	var disciplinesEntity []entity.DisciplineEntity
	for _, discipline := range disciplines {
		disciplineEntity := entity.DisciplineEntity{
			UUID:     discipline.Uuid,
			Name:     discipline.Name,
			Credits:  discipline.Credits,
			CourseID: discipline.CourseID,
		}

		disciplinesEntity = append(disciplinesEntity, disciplineEntity)
	}
	return disciplinesEntity, nil
}

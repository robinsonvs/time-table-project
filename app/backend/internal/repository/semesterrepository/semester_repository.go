package semesterrepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func (r *repository) CreateSemester(ctx context.Context, u *entity.SemesterEntity) error {
	err := r.queries.CreateSemester(ctx, sqlc.CreateSemesterParams{
		Uuid:     u.UUID,
		Semester: u.Semester,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindSemesterByID(ctx context.Context, uuid uuid.UUID) (*entity.SemesterEntity, error) {
	semester, err := r.queries.FindSemesterByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	semesterEntity := entity.SemesterEntity{
		UUID:     semester.Uuid,
		Semester: semester.Semester,
	}

	return &semesterEntity, nil
}

func (r *repository) UpdateSemester(ctx context.Context, u *entity.SemesterEntity) error {
	err := r.queries.UpdateSemester(ctx, sqlc.UpdateSemesterParams{
		Uuid:     u.UUID,
		Semester: sql.NullString{String: u.Semester, Valid: u.Semester != ""},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteSemester(ctx context.Context, uuid uuid.UUID) error {
	err := r.queries.DeleteSemester(ctx, uuid)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindManySemesters(ctx context.Context) ([]entity.SemesterEntity, error) {
	semesters, err := r.queries.FindManySemesters(ctx)
	if err != nil {
		return nil, err
	}

	var semestersEntity []entity.SemesterEntity
	for _, semester := range semesters {
		semesterEntity := entity.SemesterEntity{
			UUID:     semester.Uuid,
			Semester: semester.Semester,
		}

		semestersEntity = append(semestersEntity, semesterEntity)
	}
	return semestersEntity, nil
}

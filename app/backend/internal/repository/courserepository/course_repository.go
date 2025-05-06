package courserepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func (r *repository) CreateCourse(ctx context.Context, u *entity.CourseEntity) error {
	err := r.queries.CreateCourse(ctx, sqlc.CreateCourseParams{
		Uuid:     u.UUID,
		Name:     u.Name,
		Modality: u.Modality,
		Location: u.Location,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindCourseByID(ctx context.Context, uuid uuid.UUID) (*entity.CourseEntity, error) {
	course, err := r.queries.FindCourseByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	courseEntity := entity.CourseEntity{
		UUID:     course.Uuid,
		Name:     course.Name,
		Modality: course.Modality,
		Location: course.Location,
	}

	return &courseEntity, nil
}

func (r *repository) UpdateCourse(ctx context.Context, u *entity.CourseEntity) error {
	err := r.queries.UpdateCourse(ctx, sqlc.UpdateCourseParams{
		Uuid:     u.UUID,
		Name:     sql.NullString{String: u.Name, Valid: u.Name != ""},
		Modality: sql.NullString{String: u.Modality, Valid: u.Modality != ""},
		Location: sql.NullString{String: u.Location, Valid: u.Location != ""},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteCourse(ctx context.Context, uuid uuid.UUID) error {
	err := r.queries.DeleteCourse(ctx, uuid)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindManyCourses(ctx context.Context) ([]entity.CourseEntity, error) {
	courses, err := r.queries.FindManyCourses(ctx)
	if err != nil {
		return nil, err
	}

	var coursesEntity []entity.CourseEntity
	for _, course := range courses {
		courseEntity := entity.CourseEntity{
			ID:       course.ID,
			UUID:     course.Uuid,
			Name:     course.Name,
			Modality: course.Modality,
			Location: course.Location,
		}

		coursesEntity = append(coursesEntity, courseEntity)
	}
	return coursesEntity, nil
}

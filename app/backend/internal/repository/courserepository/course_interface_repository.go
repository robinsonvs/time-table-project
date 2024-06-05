package courserepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func NewCourseRepository(db *sql.DB, q *sqlc.Queries) CourseRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type CourseRepository interface {
	CreateCourse(ctx context.Context, u *entity.CourseEntity) error
	FindCourseByID(ctx context.Context, uuid uuid.UUID) (*entity.CourseEntity, error)
	UpdateCourse(ctx context.Context, u *entity.CourseEntity) error
	DeleteCourse(ctx context.Context, uuid uuid.UUID) error
	FindManyCourses(ctx context.Context) ([]entity.CourseEntity, error)
}

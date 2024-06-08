package semesterrepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func NewSemesterRepository(db *sql.DB, q *sqlc.Queries) SemesterRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type SemesterRepository interface {
	CreateSemester(ctx context.Context, u *entity.SemesterEntity) error
	FindSemesterByID(ctx context.Context, uuid uuid.UUID) (*entity.SemesterEntity, error)
	UpdateSemester(ctx context.Context, u *entity.SemesterEntity) error
	DeleteSemester(ctx context.Context, uuid uuid.UUID) error
	FindManySemesters(ctx context.Context) ([]entity.SemesterEntity, error)
}

package disciplinerepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func NewDisciplineRepository(db *sql.DB, q *sqlc.Queries) DisciplineRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type DisciplineRepository interface {
	CreateDiscipline(ctx context.Context, u *entity.DisciplineEntity) error
	FindDisciplineByID(ctx context.Context, uuid uuid.UUID) (*entity.DisciplineEntity, error)
	UpdateDiscipline(ctx context.Context, u *entity.DisciplineEntity) error
	DeleteDiscipline(ctx context.Context, uuid uuid.UUID) error
	FindManyDisciplines(ctx context.Context) ([]entity.DisciplineEntity, error)
	FindManyDisciplinesByCoarseId(ctx context.Context, courseId int64) ([]entity.DisciplineEntity, error)
}

package professorrepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func NewProfessorRepository(db *sql.DB, q *sqlc.Queries) ProfessorRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type ProfessorRepository interface {
	CreateProfessor(ctx context.Context, u *entity.ProfessorEntity) error
	FindProfessorByID(ctx context.Context, uuid uuid.UUID) (*entity.ProfessorEntity, error)
	UpdateProfessor(ctx context.Context, u *entity.ProfessorEntity) error
	DeleteProfessor(ctx context.Context, uuid uuid.UUID) error
	FindManyProfessors(ctx context.Context) ([]entity.ProfessorEntity, error)
}

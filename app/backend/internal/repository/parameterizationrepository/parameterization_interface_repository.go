package parameterizationrepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func NewParameterizationRepository(db *sql.DB, q *sqlc.Queries) ParameterizationRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type ParameterizationRepository interface {
	CreateParameterization(ctx context.Context, u *entity.ParameterizationEntity) error
	FindParameterizationByID(ctx context.Context, uuid uuid.UUID) (*entity.ParameterizationEntity, error)
	UpdateParameterization(ctx context.Context, u *entity.ParameterizationEntity) error
	DeleteParameterization(ctx context.Context, uuid uuid.UUID) error
	FindManyParameterizations(ctx context.Context) ([]entity.ParameterizationEntity, error)
	FindManyParameterizationsBySemesterId(ctx context.Context, semesterId int64) ([]entity.ParameterizationEntity, error)
}

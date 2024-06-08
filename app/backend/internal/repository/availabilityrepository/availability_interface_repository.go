package availabilityrepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func NewAvailabilityRepository(db *sql.DB, q *sqlc.Queries) AvailabilityRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type AvailabilityRepository interface {
	CreateAvailability(ctx context.Context, u *entity.AvailabilityEntity) error
	FindAvailabilityByID(ctx context.Context, uuid uuid.UUID) (*entity.AvailabilityEntity, error)
	UpdateAvailability(ctx context.Context, u *entity.AvailabilityEntity) error
	DeleteAvailability(ctx context.Context, uuid uuid.UUID) error
	FindManyAvailabilities(ctx context.Context) ([]entity.AvailabilityEntity, error)
	FindManyAvailabilitiesByProfessorId(ctx context.Context, professorId int64) ([]entity.AvailabilityEntity, error)
}

package eligibledisciplinerepository

import (
	"context"
	"database/sql"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func NewEligibleDisciplineRepository(db *sql.DB, q *sqlc.Queries) EligibleDisciplineRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type EligibleDisciplineRepository interface {
	CreateEligibleDiscipline(ctx context.Context, u *entity.EligibleDisciplineEntity) error
	DeleteEligibleDiscipline(ctx context.Context, u *entity.EligibleDisciplineEntity) error
}

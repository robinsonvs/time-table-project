package userrepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func NewUserRepository(db *sql.DB, q *sqlc.Queries) UserRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type UserRepository interface {
	CreateUser(ctx context.Context, u *entity.UserEntity) error
	FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	FindUserByID(ctx context.Context, uuid uuid.UUID) (*entity.UserEntity, error)
	UpdateUser(ctx context.Context, u *entity.UserEntity) error
	DeleteUser(ctx context.Context, uuid uuid.UUID) error
	FindManyUsers(ctx context.Context) ([]entity.UserEntity, error)
	UpdatePassword(ctx context.Context, pass string, uuid uuid.UUID) error
	GetUserPassword(ctx context.Context, uuid uuid.UUID) (string, error)
}

package userrepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func (r *repository) CreateUser(ctx context.Context, u *entity.UserEntity) error {
	err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Uuid:     u.UUID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	user, err := r.queries.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	userEntity := entity.UserEntity{
		ID:    user.ID,
		UUID:  user.Uuid,
		Name:  user.Name,
		Email: user.Email,
	}
	return &userEntity, nil
}

func (r *repository) FindUserByID(ctx context.Context, uuid uuid.UUID) (*entity.UserEntity, error) {
	user, err := r.queries.FindUserByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	userEntity := entity.UserEntity{
		ID:    user.ID,
		UUID:  user.Uuid,
		Name:  user.Name,
		Email: user.Email,
	}

	return &userEntity, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *entity.UserEntity) error {
	err := r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		Uuid:  u.UUID,
		Name:  sql.NullString{String: u.Name, Valid: u.Name != ""},
		Email: sql.NullString{String: u.Email, Valid: u.Email != ""},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteUser(ctx context.Context, uuid uuid.UUID) error {
	err := r.queries.DeleteUser(ctx, uuid)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindManyUsers(ctx context.Context) ([]entity.UserEntity, error) {
	users, err := r.queries.FindManyUsers(ctx)
	if err != nil {
		return nil, err
	}

	var usersEntity []entity.UserEntity
	for _, user := range users {
		userEntity := entity.UserEntity{
			ID:    user.ID,
			UUID:  user.Uuid,
			Name:  user.Name,
			Email: user.Email,
		}

		usersEntity = append(usersEntity, userEntity)
	}
	return usersEntity, nil
}

func (r *repository) UpdatePassword(ctx context.Context, pass string, uuid uuid.UUID) error {
	err := r.queries.UpdatePassword(ctx, sqlc.UpdatePasswordParams{
		Uuid:     uuid,
		Password: pass,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUserPassword(ctx context.Context, uuid uuid.UUID) (string, error) {
	pass, err := r.queries.GetUserPassword(ctx, uuid)

	if err != nil {
		return "", err
	}
	return pass, nil
}

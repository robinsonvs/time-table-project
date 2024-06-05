package userservice

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/entity"
	"github.com/robinsonvs/time-table-project/internal/handler/response"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	userExists, err := s.repo.FindUserByEmail(ctx, u.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
			return err
		}
	}

	if userExists != nil {
		slog.Error("user already exists", slog.String("package", "userservice"))
		return errors.New("user already exists")
	}

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
		return errors.New("error to encrypt password")
	}

	newUser := entity.UserEntity{
		UUID:     uuid.New(),
		Name:     u.Name,
		Email:    u.Email,
		Password: string(passwordEncrypted),
	}

	err = s.repo.CreateUser(ctx, &newUser)
	if err != nil {
		slog.Error("error to create user", "err", err, slog.String("package", "userservice"))
		return err
	}

	return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, uuid uuid.UUID) error {
	userExists, err := s.repo.FindUserByID(ctx, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("user not found", slog.String("package", "userservice"))
			return errors.New("user not found")
		}
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return err
	}

	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return errors.New("user already exists")
	}

	var updateUser entity.UserEntity
	if u.Email != "" {
		userExists, err := s.repo.FindUserByEmail(ctx, u.Email)
		if err != nil {
			if err != sql.ErrNoRows {
				slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
				return errors.New("error to search user by email")
			}
		}
		if userExists != nil {
			slog.Error("user already exists", slog.String("package", "userservice"))
			return errors.New("user already exists")
		}
		updateUser.Email = u.Email
	}

	updateUser = entity.UserEntity{
		UUID:  uuid,
		Name:  u.Name,
		Email: u.Email,
	}

	err = s.repo.UpdateUser(ctx, &updateUser)
	if err != nil {
		slog.Error("error to update user", "err", err, slog.String("package", "userservice"))
		return err
	}

	return nil
}

func (s *service) GetUserByID(ctx context.Context, uuid uuid.UUID) (*response.UserResponse, error) {
	userExists, err := s.repo.FindUserByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return nil, err
	}

	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return nil, errors.New("user not found")
	}

	user := response.UserResponse{
		UUID:  userExists.UUID.String(),
		Name:  userExists.Name,
		Email: userExists.Email,
	}

	return &user, nil
}

func (s *service) FindManyUsers(ctx context.Context) (*response.ManyUsersResponse, error) {
	findManyUsers, err := s.repo.FindManyUsers(ctx)
	if err != nil {
		slog.Error("error to find many users", "err", err, slog.String("package", "userservice"))
		return nil, err
	}

	users := response.ManyUsersResponse{}
	for _, user := range findManyUsers {
		userResponse := response.UserResponse{
			UUID:  user.UUID.String(),
			Name:  user.Name,
			Email: user.Email,
		}
		users.Users = append(users.Users, userResponse)
	}

	return &users, nil
}

func (s *service) UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, uuid uuid.UUID) error {
	userExists, err := s.repo.FindUserByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return err
	}

	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return errors.New("user not found")
	}

	oldPass, err := s.repo.GetUserPassword(ctx, uuid)
	if err != nil {
		slog.Error("error to get user password", "err", err, slog.String("package", "userservice"))
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(oldPass), []byte(u.OldPassword))
	if err != nil {
		slog.Error("invalid password", slog.String("package", "userservice"))
		return errors.New("invalid password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(oldPass), []byte(u.Password))
	if err == nil {
		slog.Error("new password is equal to old password", slog.String("package", "userservice"))
		return errors.New("new password is equal to old password")
	}

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
		return errors.New("error to encrypt password")
	}

	err = s.repo.UpdatePassword(ctx, string(passwordEncrypted), uuid)
	if err != nil {
		slog.Error("error to update password", "err", err, slog.String("package", "userservice"))
		return err
	}

	return nil
}

func (s *service) DeleteUser(ctx context.Context, uuid uuid.UUID) error {
	userExists, err := s.repo.FindUserByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return err
	}

	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return errors.New("user not found")
	}

	err = s.repo.DeleteUser(ctx, uuid)
	if err != nil {
		slog.Error("error to delete user", "err", err, slog.String("package", "userservice"))
		return err
	}

	return nil
}

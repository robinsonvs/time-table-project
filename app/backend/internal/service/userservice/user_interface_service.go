package userservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/handler/response"
	"github.com/robinsonvs/time-table-project/internal/repository/userrepository"
)

func NewUserService(repo userrepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userrepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
	UpdateUser(ctx context.Context, u dto.UpdateUserDto, uuid uuid.UUID) error
	GetUserByID(ctx context.Context, uuid uuid.UUID) (*response.UserResponse, error)
	DeleteUser(ctx context.Context, uuid uuid.UUID) error
	FindManyUsers(ctx context.Context) (*response.ManyUsersResponse, error)
	UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, uuid uuid.UUID) error
	Login(ctx context.Context, u dto.LoginDTO) (*response.UserAuthToken, error)
}

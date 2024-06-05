package courseservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/handler/response"
	"github.com/robinsonvs/time-table-project/internal/repository/courserepository"
)

func NewCourseService(repo courserepository.CourseRepository) CourseService {
	return &service{
		repo,
	}
}

type service struct {
	repo courserepository.CourseRepository
}

type CourseService interface {
	CreateCourse(ctx context.Context, u dto.CreateCourseDto) error
	UpdateCourse(ctx context.Context, u dto.UpdateCourseDto, uuid uuid.UUID) error
	GetCourseByID(ctx context.Context, uuid uuid.UUID) (*response.CourseResponse, error)
	DeleteCourse(ctx context.Context, uuid uuid.UUID) error
	FindManyCourses(ctx context.Context) (*response.ManyCoursesResponse, error)
}

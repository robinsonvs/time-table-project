package courseservice

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/entity"
	"github.com/robinsonvs/time-table-project/internal/handler/response"
	"log/slog"
)

func (s *service) CreateCourse(ctx context.Context, u dto.CreateCourseDto) error {
	//userExists, err := s.repo.FindUserByEmail(ctx, u.Email)
	//if err != nil {
	//	if err != sql.ErrNoRows {
	//		slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
	//		return err
	//	}
	//}
	//
	//if userExists != nil {
	//	slog.Error("user already exists", slog.String("package", "userservice"))
	//	return errors.New("user already exists")
	//}

	newCourse := entity.CourseEntity{
		UUID:     uuid.New(),
		Name:     u.Name,
		Modality: u.Modality,
		Location: u.Location,
	}

	err := s.repo.CreateCourse(ctx, &newCourse)
	if err != nil {
		slog.Error("error to create course", "err", err, slog.String("package", "courseservice"))
		return err
	}

	return nil
}

func (s *service) UpdateCourse(ctx context.Context, u dto.UpdateCourseDto, uuid uuid.UUID) error {
	courseExists, err := s.repo.FindCourseByID(ctx, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("course not found", slog.String("package", "courseservice"))
			return errors.New("course not found")
		}
		slog.Error("error to search course by id", "err", err, slog.String("package", "courseservice"))
		return err
	}

	if courseExists == nil {
		slog.Error("course not found", slog.String("package", "courseservice"))
		return errors.New("course already exists")
	}

	updateCourse := entity.CourseEntity{
		UUID:     uuid,
		Name:     u.Name,
		Modality: u.Modality,
		Location: u.Location,
	}

	err = s.repo.UpdateCourse(ctx, &updateCourse)
	if err != nil {
		slog.Error("error to update course", "err", err, slog.String("package", "courseservice"))
		return err
	}

	return nil
}

func (s *service) GetCourseByID(ctx context.Context, uuid uuid.UUID) (*response.CourseResponse, error) {
	courseExists, err := s.repo.FindCourseByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search course by id", "err", err, slog.String("package", "courseservice"))
		return nil, err
	}

	if courseExists == nil {
		slog.Error("course not found", slog.String("package", "courseservice"))
		return nil, errors.New("course not found")
	}

	course := response.CourseResponse{
		UUID:     courseExists.UUID.String(),
		Name:     courseExists.Name,
		Modality: courseExists.Modality,
		Location: courseExists.Location,
	}

	return &course, nil
}

func (s *service) FindManyCourses(ctx context.Context) (*response.ManyCoursesResponse, error) {
	findManyCourses, err := s.repo.FindManyCourses(ctx)
	if err != nil {
		slog.Error("error to find many courses", "err", err, slog.String("package", "courseservice"))
		return nil, err
	}

	courses := response.ManyCoursesResponse{}
	for _, course := range findManyCourses {
		courseResponse := response.CourseResponse{
			UUID:     course.UUID.String(),
			Name:     course.Name,
			Modality: course.Modality,
			Location: course.Location,
		}
		courses.Courses = append(courses.Courses, courseResponse)
	}

	return &courses, nil
}

func (s *service) DeleteCourse(ctx context.Context, uuid uuid.UUID) error {
	courseExists, err := s.repo.FindCourseByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search course id", "err", err, slog.String("package", "courseervice"))
		return err
	}

	if courseExists == nil {
		slog.Error("course not found", slog.String("package", "courseservice"))
		return errors.New("course not found")
	}

	err = s.repo.DeleteCourse(ctx, uuid)
	if err != nil {
		slog.Error("error to delete course", "err", err, slog.String("package", "courseservice"))
		return err
	}

	return nil
}

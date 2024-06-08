package semesterservice

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

func (s *service) CreateSemester(ctx context.Context, u dto.CreateSemesterDto) error {

	newSemester := entity.SemesterEntity{
		UUID:     uuid.New(),
		Semester: u.Semester,
	}

	err := s.repo.CreateSemester(ctx, &newSemester)
	if err != nil {
		slog.Error("error to create semester", "err", err, slog.String("package", "semesterservice"))
		return err
	}

	return nil
}

func (s *service) UpdateSemester(ctx context.Context, u dto.UpdateSemesterDto, uuid uuid.UUID) error {
	semesterExists, err := s.repo.FindSemesterByID(ctx, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("semester not found", slog.String("package", "semesterservice"))
			return errors.New("semester not found")
		}
		slog.Error("error to search semester by id", "err", err, slog.String("package", "semesterservice"))
		return err
	}

	if semesterExists == nil {
		slog.Error("semester not found", slog.String("package", "semesterservice"))
		return errors.New("semester already exists")
	}

	updateSemester := entity.SemesterEntity{
		UUID:     uuid,
		Semester: u.Semester,
	}

	err = s.repo.UpdateSemester(ctx, &updateSemester)
	if err != nil {
		slog.Error("error to update semester", "err", err, slog.String("package", "semesterservice"))
		return err
	}

	return nil
}

func (s *service) GetSemesterByID(ctx context.Context, uuid uuid.UUID) (*response.SemesterResponse, error) {
	semesterExists, err := s.repo.FindSemesterByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search semester by id", "err", err, slog.String("package", "semesterservice"))
		return nil, err
	}

	if semesterExists == nil {
		slog.Error("semester not found", slog.String("package", "semesterservice"))
		return nil, errors.New("semester not found")
	}

	semester := response.SemesterResponse{
		UUID:     semesterExists.UUID.String(),
		Semester: semesterExists.Semester,
	}

	return &semester, nil
}

func (s *service) FindManySemesters(ctx context.Context) (*response.ManySemestersResponse, error) {
	findManySemesters, err := s.repo.FindManySemesters(ctx)
	if err != nil {
		slog.Error("error to find many semesters", "err", err, slog.String("package", "semesterservice"))
		return nil, err
	}

	semesters := response.ManySemestersResponse{}
	for _, semesterEntity := range findManySemesters {
		semesterResponse := response.SemesterResponse{
			UUID:     semesterEntity.UUID.String(),
			Semester: semesterEntity.Semester,
		}
		semesters.Semesters = append(semesters.Semesters, semesterResponse)
	}

	return &semesters, nil
}

func (s *service) DeleteSemester(ctx context.Context, uuid uuid.UUID) error {
	semesterExists, err := s.repo.FindSemesterByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search semester id", "err", err, slog.String("package", "semesterservice"))
		return err
	}

	if semesterExists == nil {
		slog.Error("semester not found", slog.String("package", "semesterservice"))
		return errors.New("semester not found")
	}

	err = s.repo.DeleteSemester(ctx, uuid)
	if err != nil {
		slog.Error("error to delete semester", "err", err, slog.String("package", "semesterservice"))
		return err
	}

	return nil
}

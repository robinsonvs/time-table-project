package disciplineservice

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

func (s *service) CreateDiscipline(ctx context.Context, u dto.CreateDisciplineDto) error {

	newDiscipline := entity.DisciplineEntity{
		UUID:     uuid.New(),
		Name:     u.Name,
		Credits:  u.Credits,
		CourseID: u.CourseId,
	}

	err := s.repo.CreateDiscipline(ctx, &newDiscipline)
	if err != nil {
		slog.Error("error to create discipline", "err", err, slog.String("package", "disciplineservice"))
		return err
	}

	return nil
}

func (s *service) UpdateDiscipline(ctx context.Context, u dto.UpdateDisciplineDto, uuid uuid.UUID) error {
	disciplineExists, err := s.repo.FindDisciplineByID(ctx, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("discipline not found", slog.String("package", "disciplineservice"))
			return errors.New("discipline not found")
		}
		slog.Error("error to search discipline by id", "err", err, slog.String("package", "disciplineservice"))
		return err
	}

	if disciplineExists == nil {
		slog.Error("discipline not found", slog.String("package", "disciplineservice"))
		return errors.New("discipline already exists")
	}

	updateDiscipline := entity.DisciplineEntity{
		UUID:    uuid,
		Name:    u.Name,
		Credits: u.Credits,
	}

	err = s.repo.UpdateDiscipline(ctx, &updateDiscipline)
	if err != nil {
		slog.Error("error to update discipline", "err", err, slog.String("package", "disciplineservice"))
		return err
	}

	return nil
}

func (s *service) GetDisciplineByID(ctx context.Context, uuid uuid.UUID) (*response.DisciplineResponse, error) {
	disciplineExists, err := s.repo.FindDisciplineByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search discipline by id", "err", err, slog.String("package", "disciplineservice"))
		return nil, err
	}

	if disciplineExists == nil {
		slog.Error("discipline not found", slog.String("package", "disciplineservice"))
		return nil, errors.New("discipline not found")
	}

	discipline := response.DisciplineResponse{
		UUID:     disciplineExists.UUID.String(),
		Name:     disciplineExists.Name,
		Credits:  disciplineExists.Credits,
		CourseId: disciplineExists.CourseID,
	}

	return &discipline, nil
}

func (s *service) FindManyDisciplines(ctx context.Context) (*response.ManyDisciplinesResponse, error) {
	findManyDisciplines, err := s.repo.FindManyDisciplines(ctx)
	if err != nil {
		slog.Error("error to find many disciplines", "err", err, slog.String("package", "disciplineservice"))
		return nil, err
	}

	disciplines := response.ManyDisciplinesResponse{}
	for _, disciplineEntity := range findManyDisciplines {
		disciplineResponse := response.DisciplineResponse{
			Id:       disciplineEntity.ID,
			UUID:     disciplineEntity.UUID.String(),
			Name:     disciplineEntity.Name,
			Credits:  disciplineEntity.Credits,
			CourseId: disciplineEntity.CourseID,
		}
		disciplines.Disciplines = append(disciplines.Disciplines, disciplineResponse)
	}

	return &disciplines, nil
}

func (s *service) FindManyDisciplinesByCourseId(ctx context.Context, courseId int64) (*response.ManyDisciplinesResponse, error) {
	findManyDisciplinesByCourse, err := s.repo.FindManyDisciplinesByCoarseId(ctx, courseId)
	if err != nil {
		slog.Error("error to find many disciplines", "err", err, slog.String("package", "disciplineservice"))
		return nil, err
	}

	disciplinesByCourse := response.ManyDisciplinesResponse{}
	for _, disciplineEntity := range findManyDisciplinesByCourse {
		disciplineResponse := response.DisciplineResponse{
			Id:       disciplineEntity.ID,
			UUID:     disciplineEntity.UUID.String(),
			Name:     disciplineEntity.Name,
			Credits:  disciplineEntity.Credits,
			CourseId: disciplineEntity.CourseID,
		}
		disciplinesByCourse.Disciplines = append(disciplinesByCourse.Disciplines, disciplineResponse)
	}

	return &disciplinesByCourse, nil
}

func (s *service) DeleteDiscipline(ctx context.Context, uuid uuid.UUID) error {
	disciplineExists, err := s.repo.FindDisciplineByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search discipline id", "err", err, slog.String("package", "disciplineservice"))
		return err
	}

	if disciplineExists == nil {
		slog.Error("discipline not found", slog.String("package", "disciplineservice"))
		return errors.New("discipline not found")
	}

	err = s.repo.DeleteDiscipline(ctx, uuid)
	if err != nil {
		slog.Error("error to delete discipline", "err", err, slog.String("package", "disciplineservice"))
		return err
	}

	return nil
}

package professorservice

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

func (s *service) CreateProfessor(ctx context.Context, u dto.CreateProfessorDto) error {

	newProfessor := entity.ProfessorEntity{
		UUID:            uuid.New(),
		Name:            u.Name,
		HoursToAllocate: u.HoursToAllocate,
	}

	err := s.repo.CreateProfessor(ctx, &newProfessor)
	if err != nil {
		slog.Error("error to create professor", "err", err, slog.String("package", "professorservice"))
		return err
	}

	return nil
}

func (s *service) UpdateProfessor(ctx context.Context, u dto.UpdateProfessorDto, uuid uuid.UUID) error {
	professorExists, err := s.repo.FindProfessorByID(ctx, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("professor not found", slog.String("package", "professorservice"))
			return errors.New("professor not found")
		}
		slog.Error("error to search professor by id", "err", err, slog.String("package", "professorservice"))
		return err
	}

	if professorExists == nil {
		slog.Error("professor not found", slog.String("package", "professorservice"))
		return errors.New("professor already exists")
	}

	updateProfessor := entity.ProfessorEntity{
		UUID:            uuid,
		Name:            u.Name,
		HoursToAllocate: u.HoursToAllocate,
	}

	err = s.repo.UpdateProfessor(ctx, &updateProfessor)
	if err != nil {
		slog.Error("error to update professor", "err", err, slog.String("package", "professorservice"))
		return err
	}

	return nil
}

func (s *service) GetProfessorByID(ctx context.Context, uuid uuid.UUID) (*response.ProfessorResponse, error) {
	professorExists, err := s.repo.FindProfessorByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search professor by id", "err", err, slog.String("package", "professorservice"))
		return nil, err
	}

	if professorExists == nil {
		slog.Error("professor not found", slog.String("package", "professorservice"))
		return nil, errors.New("professor not found")
	}

	professor := response.ProfessorResponse{
		UUID:            professorExists.UUID.String(),
		Name:            professorExists.Name,
		HoursToAllocate: professorExists.HoursToAllocate,
	}

	return &professor, nil
}

func (s *service) FindManyProfessors(ctx context.Context) (*response.ManyProfessorsResponse, error) {
	findManyProfessors, err := s.repo.FindManyProfessors(ctx)
	if err != nil {
		slog.Error("error to find many professors", "err", err, slog.String("package", "professorservice"))
		return nil, err
	}

	professors := response.ManyProfessorsResponse{}
	for _, professorEntity := range findManyProfessors {
		professorResponse := response.ProfessorResponse{
			UUID:            professorEntity.UUID.String(),
			Name:            professorEntity.Name,
			HoursToAllocate: professorEntity.HoursToAllocate,
		}
		professors.Professors = append(professors.Professors, professorResponse)
	}

	return &professors, nil
}

func (s *service) DeleteProfessor(ctx context.Context, uuid uuid.UUID) error {
	professorExists, err := s.repo.FindProfessorByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search professor id", "err", err, slog.String("package", "professorservice"))
		return err
	}

	if professorExists == nil {
		slog.Error("professor not found", slog.String("package", "professorservice"))
		return errors.New("professor not found")
	}

	err = s.repo.DeleteProfessor(ctx, uuid)
	if err != nil {
		slog.Error("error to delete professor", "err", err, slog.String("package", "professorservice"))
		return err
	}

	return nil
}

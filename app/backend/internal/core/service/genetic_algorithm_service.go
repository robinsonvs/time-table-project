package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/core/process"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func (s *GeneticAlgorithmService) GenerateProposal(ctx context.Context, parameterizationID uuid.UUID) error {
	parameterization, err := s.ParameterizationRepo.FindParameterizationByID(ctx, parameterizationID)
	if err != nil {
		return err
	}

	parameterization.Disciplines, err = s.ParameterizationRepo.GetDisciplinesByCourseID(ctx, parameterization.CourseID)
	if err != nil {
		return err
	}

	parameterization.Professors, err = s.ParameterizationRepo.GetProfessorsByCourseID(ctx, parameterization.CourseID)
	if err != nil {
		return err
	}

	disciplines, err := s.DisciplineRepo.FindManyDisciplinesByCoarseId(ctx, parameterization.CourseID)
	if err != nil {
		return err
	}

	professors, err := s.ProfessorRepo.GetProfessorsWithDisciplines(ctx)
	if err != nil {
		return err
	}

	availabilities, err := s.AvailabilityRepo.FindManyAvailabilities(ctx)
	if err != nil {
		return err
	}

	bestTimetable := process.RunGeneticAlgorithm(disciplines, professors, availabilities, *parameterization, 1)

	proposal := &entity.ProposalEntity{
		SemesterID: parameterization.SemesterID,
		CourseID:   parameterization.CourseID,
		Classes:    bestTimetable.Classes,
	}

	err = s.ParameterizationRepo.CreateProposal(ctx, proposal)
	if err != nil {
		return err
	}

	return nil
}

package parameterizationservice

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

func (s *service) CreateParameterization(ctx context.Context, u dto.CreateParameterizationDto) error {

	newParameterization := entity.ParameterizationEntity{
		UUID:                    uuid.New(),
		MaxCreditsToOffer:       u.MaxCreditsToOffer,
		NumClassesPerDiscipline: u.NumClassesPerDiscipline,
		SemesterID:              u.SemesterId,
		CourseID:                u.CourseId,
	}

	err := s.repo.CreateParameterization(ctx, &newParameterization)
	if err != nil {
		slog.Error("error to create parameterization", "err", err, slog.String("package", "parameterizationservice"))
		return err
	}

	return nil
}

func (s *service) UpdateParameterization(ctx context.Context, u dto.UpdateParameterizationDto, uuid uuid.UUID) error {
	parameterizationExists, err := s.repo.FindParameterizationByID(ctx, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("parameterization not found", slog.String("package", "parameterizationservice"))
			return errors.New("parameterization not found")
		}
		slog.Error("error to search parameterization by id", "err", err, slog.String("package", "parameterizationservice"))
		return err
	}

	if parameterizationExists == nil {
		slog.Error("parameterization not found", slog.String("package", "parameterizationservice"))
		return errors.New("parameterization already exists")
	}

	updateParameterization := entity.ParameterizationEntity{
		UUID:                    uuid,
		MaxCreditsToOffer:       u.MaxCreditsToOffer,
		NumClassesPerDiscipline: u.NumClassesPerDiscipline,
		SemesterID:              u.SemesterId,
		CourseID:                u.CourseId,
	}

	err = s.repo.UpdateParameterization(ctx, &updateParameterization)
	if err != nil {
		slog.Error("error to update discipline", "err", err, slog.String("package", "disciplineservice"))
		return err
	}

	return nil
}

func (s *service) GetParameterizationByID(ctx context.Context, uuid uuid.UUID) (*response.ParameterizationResponse, error) {
	parameterizationExists, err := s.repo.FindParameterizationByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search parameterization by id", "err", err, slog.String("package", "parameterizationservice"))
		return nil, err
	}

	if parameterizationExists == nil {
		slog.Error("parameterization not found", slog.String("package", "parameterizationservice"))
		return nil, errors.New("parameterization not found")
	}

	parameterization := response.ParameterizationResponse{
		UUID:                    parameterizationExists.UUID.String(),
		MaxCreditsToOffer:       parameterizationExists.MaxCreditsToOffer,
		NumClassesPerDiscipline: parameterizationExists.NumClassesPerDiscipline,
		SemesterId:              parameterizationExists.SemesterID,
		CourseId:                parameterizationExists.CourseID,
	}

	return &parameterization, nil
}

func (s *service) FindManyParameterizations(ctx context.Context) (*response.ManyParameterizationsResponse, error) {
	findManyParameterizations, err := s.repo.FindManyParameterizations(ctx)
	if err != nil {
		slog.Error("error to find many parameterizations", "err", err, slog.String("package", "parameterizationsservice"))
		return nil, err
	}

	parameterizations := response.ManyParameterizationsResponse{}
	for _, parameterizationEntity := range findManyParameterizations {
		parameterizationResponse := response.ParameterizationResponse{
			UUID:                    parameterizationEntity.UUID.String(),
			MaxCreditsToOffer:       parameterizationEntity.MaxCreditsToOffer,
			NumClassesPerDiscipline: parameterizationEntity.NumClassesPerDiscipline,
			SemesterId:              parameterizationEntity.SemesterID,
			CourseId:                parameterizationEntity.CourseID,
		}
		parameterizations.Parameterizations = append(parameterizations.Parameterizations, parameterizationResponse)
	}

	return &parameterizations, nil
}

func (s *service) FindManyParameterizationsBySemesterId(ctx context.Context, courseId int64) (*response.ManyParameterizationsResponse, error) {
	findManyParameterizationsBySemester, err := s.repo.FindManyParameterizationsBySemesterId(ctx, courseId)
	if err != nil {
		slog.Error("error to find many parameterizations", "err", err, slog.String("package", "parameterizationservice"))
		return nil, err
	}

	parameterizationsBySemester := response.ManyParameterizationsResponse{}
	for _, parameterizationEntity := range findManyParameterizationsBySemester {
		parameterizationResponse := response.ParameterizationResponse{
			UUID:                    parameterizationEntity.UUID.String(),
			MaxCreditsToOffer:       parameterizationEntity.MaxCreditsToOffer,
			NumClassesPerDiscipline: parameterizationEntity.NumClassesPerDiscipline,
			SemesterId:              parameterizationEntity.SemesterID,
			CourseId:                parameterizationEntity.CourseID,
		}
		parameterizationsBySemester.Parameterizations = append(parameterizationsBySemester.Parameterizations, parameterizationResponse)
	}

	return &parameterizationsBySemester, nil
}

func (s *service) DeleteParameterization(ctx context.Context, uuid uuid.UUID) error {
	parameterizationExists, err := s.repo.FindParameterizationByID(ctx, uuid)
	if err != nil {
		slog.Error("error to search parameterization id", "err", err, slog.String("package", "parameterizationservice"))
		return err
	}

	if parameterizationExists == nil {
		slog.Error("parameterization not found", slog.String("package", "parameterizationservice"))
		return errors.New("parameterization not found")
	}

	err = s.repo.DeleteParameterization(ctx, uuid)
	if err != nil {
		slog.Error("error to delete parameterization", "err", err, slog.String("package", "parameterizationservice"))
		return err
	}

	return nil
}

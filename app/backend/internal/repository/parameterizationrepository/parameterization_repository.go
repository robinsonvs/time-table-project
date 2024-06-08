package parameterizationrepository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/entity"
)

func (r *repository) CreateParameterization(ctx context.Context, u *entity.ParameterizationEntity) error {
	err := r.queries.CreateParameterization(ctx, sqlc.CreateParameterizationParams{
		Uuid:                    u.UUID,
		Maxcreditstooffer:       u.MaxCreditsToOffer,
		Numclassesperdiscipline: u.NumClassesPerDiscipline,
		SemesterID:              u.SemesterID,
		CourseID:                u.CourseID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindParameterizationByID(ctx context.Context, uuid uuid.UUID) (*entity.ParameterizationEntity, error) {
	parameterization, err := r.queries.FindParameterizationByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	parameterizationEntity := entity.ParameterizationEntity{
		UUID:                    parameterization.Uuid,
		MaxCreditsToOffer:       parameterization.Maxcreditstooffer,
		NumClassesPerDiscipline: parameterization.Numclassesperdiscipline,
		SemesterID:              parameterization.SemesterID,
		CourseID:                parameterization.CourseID,
	}

	return &parameterizationEntity, nil
}

func (r *repository) UpdateParameterization(ctx context.Context, u *entity.ParameterizationEntity) error {
	err := r.queries.UpdateParameterization(ctx, sqlc.UpdateParameterizationParams{
		Uuid:                    u.UUID,
		MaxCreditsToOffer:       sql.NullInt32{Int32: u.MaxCreditsToOffer, Valid: u.MaxCreditsToOffer != 0},
		NumClassesPerDiscipline: sql.NullInt32{Int32: u.NumClassesPerDiscipline, Valid: u.NumClassesPerDiscipline != 0},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteParameterization(ctx context.Context, uuid uuid.UUID) error {
	err := r.queries.DeleteParameterization(ctx, uuid)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindManyParameterizations(ctx context.Context) ([]entity.ParameterizationEntity, error) {
	parameterizations, err := r.queries.FindManyParameterizations(ctx)
	if err != nil {
		return nil, err
	}

	var parameterizationsEntity []entity.ParameterizationEntity
	for _, parameterization := range parameterizations {
		parameterizationEntity := entity.ParameterizationEntity{
			UUID:                    parameterization.Uuid,
			MaxCreditsToOffer:       parameterization.Maxcreditstooffer,
			NumClassesPerDiscipline: parameterization.Numclassesperdiscipline,
			SemesterID:              parameterization.SemesterID,
			CourseID:                parameterization.CourseID,
		}

		parameterizationsEntity = append(parameterizationsEntity, parameterizationEntity)
	}
	return parameterizationsEntity, nil
}

func (r *repository) FindManyParameterizationsBySemesterId(ctx context.Context, semesterId int64) ([]entity.ParameterizationEntity, error) {
	parameterizations, err := r.queries.FindManyParameterizationsBySemesterId(ctx, semesterId)
	if err != nil {
		return nil, err
	}

	var parameterizationsEntity []entity.ParameterizationEntity
	for _, parameterization := range parameterizations {
		parameterizationEntity := entity.ParameterizationEntity{
			UUID:                    parameterization.Uuid,
			MaxCreditsToOffer:       parameterization.Maxcreditstooffer,
			NumClassesPerDiscipline: parameterization.Numclassesperdiscipline,
			SemesterID:              parameterization.SemesterID,
			CourseID:                parameterization.CourseID,
		}

		parameterizationsEntity = append(parameterizationsEntity, parameterizationEntity)
	}
	return parameterizationsEntity, nil
}

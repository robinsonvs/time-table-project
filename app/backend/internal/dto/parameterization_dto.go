package dto

type CreateParameterizationDto struct {
	MaxCreditsToOffer       int32 `json:"maxCreditsToOffer" validate:"required"`
	NumClassesPerDiscipline int32 `json:"numClassesPerDiscipline" validate:"required"`
	SemesterId              int64 `json:"semester_id" validate:"required"`
	CourseId                int64 `json:"course_id" validate:"required"`
}

type UpdateParameterizationDto struct {
	MaxCreditsToOffer       int32 `json:"maxCreditsToOffer" validate:"required"`
	NumClassesPerDiscipline int32 `json:"numClassesPerDiscipline" validate:"required"`
}

package dto

type CreateParameterizationDto struct {
	MaxCreditsToOffer       int32 `json:"maxCreditsToOffer" validate:"required,min=3,max=3"`
	NumClassesPerDiscipline int32 `json:"numClassesPerDiscipline" validate:"required,min=3,max=3"`
	SemesterId              int64 `json:"semester_id" validate:"required,min=1,max=100"`
	CourseId                int64 `json:"course_id" validate:"required,min=1,max=100"`
}

type UpdateParameterizationDto struct {
	MaxCreditsToOffer       int32 `json:"maxCreditsToOffer" validate:"required,min=3,max=3"`
	NumClassesPerDiscipline int32 `json:"numClassesPerDiscipline" validate:"required,min=3,max=3"`
	SemesterId              int64 `json:"semester_id" validate:"required,min=1,max=100"`
	CourseId                int64 `json:"course_id" validate:"required,min=1,max=100"`
}

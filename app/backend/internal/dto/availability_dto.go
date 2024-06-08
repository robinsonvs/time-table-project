package dto

type CreateAvailabilityDto struct {
	DayOfWeek   string `json:"dayOfWeek" validate:"required,min=3,max=255"`
	Shift       string `json:"shift" validate:"required,min=3,max=255"`
	ProfessorId int64  `json:"professor_id" validate:"required,min=1,max=100"`
}

type UpdateAvailabilityDto struct {
	DayOfWeek   string `json:"dayOfWeek" validate:"required,min=3,max=255"`
	Shift       string `json:"shift" validate:"required,min=3,max=255"`
	ProfessorId int64  `json:"professor_id" validate:"required,min=1,max=100"`
}

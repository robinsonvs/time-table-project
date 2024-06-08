package dto

type CreateSemesterDto struct {
	Semester string `json:"semester" validate:"required,min=3,max=255"`
}

type UpdateSemesterDto struct {
	Semester string `json:"semester" validate:"omitempty,min=3,max=255"`
}

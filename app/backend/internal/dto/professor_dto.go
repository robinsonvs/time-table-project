package dto

type CreateProfessorDto struct {
	Name            string `json:"name" validate:"required,min=3,max=255"`
	HoursToAllocate int32  `json:"hoursToAllocate" validate:"required,min=1,max=3"`
}

type UpdateProfessorDto struct {
	Name            string `json:"name" validate:"required,min=3,max=255"`
	HoursToAllocate int32  `json:"hoursToAllocate" validate:"required,min=1,max=3"`
}

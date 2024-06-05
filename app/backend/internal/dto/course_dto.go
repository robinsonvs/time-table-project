package dto

type CreateCourseDto struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Modality string `json:"modality" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type UpdateCourseDto struct {
	Name     string `json:"name" validate:"omitempty,min=3,max=255"`
	Modality string `json:"modality" validate:"omitempty"`
	Location string `json:"location" validate:"omitempty"`
}

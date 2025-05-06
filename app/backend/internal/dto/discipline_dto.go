package dto

type CreateDisciplineDto struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Credits  int32  `json:"credits" validate:"required"`
	CourseId int64  `json:"course_id" validate:"required"`
}

type UpdateDisciplineDto struct {
	Name    string `json:"name" validate:"required,min=3,max=255"`
	Credits int32  `json:"credits" validate:"required"`
}

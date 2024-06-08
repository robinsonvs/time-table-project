package dto

type CreateDisciplineDto struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Credits  int32  `json:"credits" validate:"required,min=1,max=100"`
	CourseId int64  `json:"course_id" validate:"required,min=1,max=100"`
}

type UpdateDisciplineDto struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Credits  int32  `json:"credits" validate:"required,min=1,max=100"`
	CourseId int64  `json:"course_id" validate:"required,min=1,max=100"`
}

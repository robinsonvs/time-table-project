package dto

type CreateEligibleDisciplineDto struct {
	ProfessorId  int64 `json:"professor_id" validate:"required"`
	DisciplineId int64 `json:"discipline_id" validate:"required"`
}

type DeleteEligibleDisciplineDto struct {
	ProfessorId  int64 `json:"professor_id" validate:"required"`
	DisciplineId int64 `json:"discipline_id" validate:"required"`
}

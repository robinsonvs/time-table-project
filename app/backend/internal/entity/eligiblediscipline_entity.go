package entity

type EligibleDisciplineEntity struct {
	ID           int64 `json:"id"`
	ProfessorID  int64 `json:"professor_id"`
	DisciplineID int64 `json:"discipline_id"`
}

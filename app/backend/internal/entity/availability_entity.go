package entity

type Availability struct {
	ID          int64  `json:"id"`
	UUID        string `json:"uuid"`
	DayOfWeek   string `json:"day_of_week"`
	Shift       string `json:"shift"`
	ProfessorID int64  `json:"professor_id"`
}

package entity

import "github.com/google/uuid"

type AvailabilityEntity struct {
	ID          int64     `json:"id"`
	UUID        uuid.UUID `json:"uuid"`
	DayOfWeek   string    `json:"day_of_week"`
	Shift       string    `json:"shift"`
	ProfessorID int64     `json:"professor_id"`
}

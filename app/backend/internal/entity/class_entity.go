package entity

import (
	"github.com/google/uuid"
	"time"
)

type ClassEntity struct {
	ID           int64     `json:"id"`
	UUID         uuid.UUID `json:"uuid"`
	DayOfWeek    string    `json:"day_of_week"`
	Shift        string    `json:"shift"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	DisciplineID int64     `json:"discipline_id"`
	ProfessorID  int64     `json:"professor_id"`
	ProposalID   int64     `json:"proposal_id"`
}

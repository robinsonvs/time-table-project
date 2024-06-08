package entity

import "github.com/google/uuid"

type ProfessorEntity struct {
	ID              int64     `json:"id"`
	UUID            uuid.UUID `json:"uuid"`
	Name            string    `json:"name"`
	HoursToAllocate int32     `json:"hoursToAllocate"`
}

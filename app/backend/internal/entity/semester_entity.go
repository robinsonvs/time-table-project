package entity

import "github.com/google/uuid"

type SemesterEntity struct {
	ID       int64     `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	Semester string    `json:"semester"`
}

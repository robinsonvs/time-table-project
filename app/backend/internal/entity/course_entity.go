package entity

import "github.com/google/uuid"

type CourseEntity struct {
	ID       int64     `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Modality string    `json:"modality"`
	Location string    `json:"location"`
}

package entity

import "github.com/google/uuid"

type DisciplineEntity struct {
	ID       int64     `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Credits  int32     `json:"credits"`
	CourseID int64     `json:"course_id"`
}

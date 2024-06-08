package entity

import "github.com/google/uuid"

type ProposalEntity struct {
	ID         int64         `json:"id"`
	UUID       uuid.UUID     `json:"uuid"`
	SemesterID int64         `json:"semester_id"`
	CourseID   int64         `json:"course_id"`
	Classes    []ClassEntity `json:"classes"`
}

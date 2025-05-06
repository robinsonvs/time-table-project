package entity

import "github.com/google/uuid"

type ParameterizationEntity struct {
	ID                      int64              `json:"id"`
	UUID                    uuid.UUID          `json:"uuid"`
	MaxCreditsToOffer       int32              `json:"max_credits_to_offer"`
	NumClassesPerDiscipline int32              `json:"num_classes_per_discipline"`
	SemesterID              int64              `json:"semester_id"`
	CourseID                int64              `json:"course_id"`
	Disciplines             []DisciplineEntity `json:"disciplines"`
	Professors              []ProfessorEntity  `json:"professors"`
}

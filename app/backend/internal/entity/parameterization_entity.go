package entity

type Parameterization struct {
	ID                      int64  `json:"id"`
	UUID                    string `json:"uuid"`
	MaxCreditsToOffer       int    `json:"max_credits_to_offer"`
	NumClassesPerDiscipline int    `json:"num_classes_per_discipline"`
	SemesterID              int64  `json:"semester_id"`
	CourseID                int64  `json:"course_id"`
}

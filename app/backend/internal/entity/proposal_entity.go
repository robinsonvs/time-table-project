package entity

type Proposal struct {
	ID         int64  `json:"id"`
	UUID       string `json:"uuid"`
	SemesterID int64  `json:"semester_id"`
	CourseID   int64  `json:"course_id"`
}

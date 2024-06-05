package entity

type Discipline struct {
	ID       int64  `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Credits  int    `json:"credits"`
	CourseID int64  `json:"course_id"`
}

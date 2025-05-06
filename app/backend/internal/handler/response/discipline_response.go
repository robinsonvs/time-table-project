package response

type DisciplineResponse struct {
	Id       int64  `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Credits  int32  `json:"credits"`
	CourseId int64  `json:"course_id"`
}

type ManyDisciplinesResponse struct {
	Disciplines []DisciplineResponse `json:"disciplines"`
}

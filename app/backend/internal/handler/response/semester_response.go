package response

type SemesterResponse struct {
	Id       int64  `json:"id"`
	UUID     string `json:"uuid"`
	Semester string `json:"semester"`
}

type ManySemestersResponse struct {
	Semesters []SemesterResponse `json:"semesters"`
}

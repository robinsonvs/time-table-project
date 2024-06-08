package response

type SemesterResponse struct {
	UUID     string `json:"uuid"`
	Semester string `json:"semester"`
}

type ManySemestersResponse struct {
	Semesters []SemesterResponse `json:"semesters"`
}

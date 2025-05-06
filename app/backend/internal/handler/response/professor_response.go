package response

type ProfessorResponse struct {
	Id              int64  `json:"id"`
	UUID            string `json:"uuid"`
	Name            string `json:"name"`
	HoursToAllocate int32  `json:"hoursToAllocate"`
}

type ManyProfessorsResponse struct {
	Professors []ProfessorResponse `json:"professors"`
}

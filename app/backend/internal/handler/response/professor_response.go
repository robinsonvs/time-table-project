package response

type ProfessorResponse struct {
	UUID            string `json:"uuid"`
	Name            string `json:"name"`
	HoursToAllocate int32  `json:"hoursToAllocate"`
}

type ManyProfessorsResponse struct {
	Professors []ProfessorResponse `json:"professors"`
}

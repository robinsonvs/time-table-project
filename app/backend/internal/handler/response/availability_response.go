package response

type AvailabilityResponse struct {
	Id          int64  `json:"'id'"`
	UUID        string `json:"uuid"`
	DayOfWeek   string `json:"dayOfWeek"`
	Shift       string `json:"shift"`
	ProfessorId int64  `json:"professor_id"`
}

type ManyAvailabilitiesResponse struct {
	Availabilities []AvailabilityResponse `json:"availabilities"`
}

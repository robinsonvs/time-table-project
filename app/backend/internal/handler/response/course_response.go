package response

type CourseResponse struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Modality string `json:"modality"`
	Location string `json:"location"`
}

type ManyCoursesResponse struct {
	Courses []CourseResponse `json:"courses"`
}

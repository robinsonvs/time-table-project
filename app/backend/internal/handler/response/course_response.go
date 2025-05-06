package response

type CourseResponse struct {
	Id       int64  `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Modality string `json:"modality"`
	Location string `json:"location"`
}

type ManyCoursesResponse struct {
	Courses []CourseResponse `json:"courses"`
}

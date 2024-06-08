package response

type ParameterizationResponse struct {
	UUID                    string `json:"uuid"`
	MaxCreditsToOffer       int32  `json:"maxCreditsToOffer"`
	NumClassesPerDiscipline int32  `json:"numClassesPerDiscipline"`
	SemesterId              int64  `json:"semester_id"`
	CourseId                int64  `json:"course_id"`
}

type ManyParameterizationsResponse struct {
	Parameterizations []ParameterizationResponse `json:"parameterizations"`
}

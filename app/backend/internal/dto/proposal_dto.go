package dto

import (
	"time"
)

type DisciplineDTO struct {
	ID       int64  `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Credits  int    `json:"credits"`
	CourseID int64  `json:"course_id"`
}

type ProfessorDTO struct {
	ID              int64           `json:"id"`
	UUID            string          `json:"uuid"`
	Name            string          `json:"name"`
	HoursToAllocate int             `json:"hours_to_allocate"`
	Disciplines     []DisciplineDTO `json:"disciplines"`
}

type ClassDTO struct {
	ID         int64         `json:"id"`
	UUID       string        `json:"uuid"`
	DayOfWeek  string        `json:"day_of_week"`
	Shift      string        `json:"shift"`
	StartTime  time.Time     `json:"start_time"`
	EndTime    time.Time     `json:"end_time"`
	Discipline DisciplineDTO `json:"discipline"`
	Professor  ProfessorDTO  `json:"professor"`
	ProposalID int64         `json:"proposal_id"`
}

type ProposalDTO struct {
	ID         int64      `json:"id"`
	UUID       string     `json:"uuid"`
	SemesterID int64      `json:"semester_id"`
	CourseID   int64      `json:"course_id"`
	Classes    []ClassDTO `json:"classes"`
}

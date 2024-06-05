package handler

import (
	"github.com/robinsonvs/time-table-project/internal/service/courseservice"
	"github.com/robinsonvs/time-table-project/internal/service/userservice"
	"net/http"
)

func NewHandler(userService userservice.UserService,
	courseService courseservice.CourseService) Handler {
	return &handler{
		userService:   userService,
		courseService: courseService,
	}
}

type handler struct {
	userService   userservice.UserService
	courseService courseservice.CourseService
}

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindManyUsers(w http.ResponseWriter, r *http.Request)
	
	UpdateUserPassword(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)

	CreateCourse(w http.ResponseWriter, r *http.Request)
	UpdateCourse(w http.ResponseWriter, r *http.Request)
	DeleteCourse(w http.ResponseWriter, r *http.Request)
	GetCourseByID(w http.ResponseWriter, r *http.Request)
	FindManyCourses(w http.ResponseWriter, r *http.Request)
}

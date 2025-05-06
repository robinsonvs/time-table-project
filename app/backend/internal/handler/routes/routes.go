package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/robinsonvs/time-table-project/config/env"
	"github.com/robinsonvs/time-table-project/internal/handler"
	"github.com/robinsonvs/time-table-project/internal/handler/middleware"
)

func InitRoutes(router chi.Router, h handler.Handler) {
	router.Use(middleware.LoggerData)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
	})

	router.Post("/users", h.CreateUser)

	router.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(env.Env.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Patch("/users", h.UpdateUser)
		r.Get("/users/{uuid}", h.GetUserByID)
		r.Delete("/users/{uuid}", h.DeleteUser)
		r.Get("/users/list-all", h.FindManyUsers)
		r.Patch("/users/password", h.UpdateUserPassword)

		r.Post("/courses", h.CreateCourse)
		r.Patch("/courses/{uuid}", h.UpdateCourse)
		r.Delete("/courses/{uuid}", h.DeleteCourse)
		r.Get("/courses/{uuid}", h.GetCourseByID)
		r.Get("/courses/list-all", h.FindManyCourses)

		r.Post("/semesters", h.CreateSemester)
		r.Patch("/semesters/{uuid}", h.UpdateSemester)
		r.Delete("/semesters/{uuid}", h.DeleteSemester)
		r.Get("/semesters/{uuid}", h.GetSemesterByID)
		r.Get("/semesters/list-all", h.FindManySemesters)

		r.Post("/professors", h.CreateProfessor)
		r.Patch("/professors/{uuid}", h.UpdateProfessor)
		r.Delete("/professors/{uuid}", h.DeleteProfessor)
		r.Get("/professors/{uuid}", h.GetProfessorByID)
		r.Get("/professors/list-all", h.FindManyProfessors)

		r.Post("/disciplines", h.CreateDiscipline)
		r.Patch("/disciplines/{uuid}", h.UpdateDiscipline)
		r.Delete("/disciplines/{uuid}", h.DeleteDiscipline)
		r.Get("/disciplines/{uuid}", h.GetDisciplineByID)
		r.Get("/disciplines/list-all", h.FindManyDisciplines)
		r.Get("/disciplines/list-all/{courseId}", h.FindManyDisciplinesByCourseId)

		r.Post("/availabilities", h.CreateAvailability)
		r.Patch("/availabilities/{uuid}", h.UpdateAvailability)
		r.Delete("/availabilities/{uuid}", h.DeleteAvailability)
		r.Get("/availabilities/{uuid}", h.GetAvailabilityByID)
		r.Get("/availabilities/list-all", h.FindManyAvailabilities)
		r.Get("/availabilities/list-all/{professorId}", h.FindManyAvailabilitiesByProfessorId)

		r.Post("/parameterizations", h.CreateParameterization)
		r.Patch("/parameterizations/{uuid}", h.UpdateParameterization)
		r.Delete("/parameterizations/{uuid}", h.DeleteParameterization)
		r.Get("/parameterizations/{uuid}", h.GetParameterizationByID)
		r.Get("/parameterizations/list-all", h.FindManyParameterizations)
		r.Get("/parameterizations/list-all/{semesterId}", h.FindManyParameterizationsBySemesterId)

		r.Post("/eligible-disciplines", h.CreateEligibleDiscipline)
		r.Delete("/eligible-disciplines", h.DeleteEligibleDiscipline)

		r.Post("/generate-proposal/{parameterizationID}", h.GenerateProposal)

	})

}

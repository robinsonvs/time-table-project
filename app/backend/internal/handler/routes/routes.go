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
		r.Patch("/users/{uuid}/password", h.UpdateUserPassword)

		r.Post("/courses", h.CreateCourse)
		r.Patch("/courses/{uuid}", h.UpdateCourse)
		r.Delete("/courses/{uuid}", h.DeleteCourse)
		r.Get("/courses/{uuid}", h.GetCourseByID)
		r.Get("/courses/list-all", h.FindManyCourses)

	})

}

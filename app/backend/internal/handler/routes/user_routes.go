package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/robinsonvs/time-table-project/config/env"
	"github.com/robinsonvs/time-table-project/internal/handler/middleware"
	"github.com/robinsonvs/time-table-project/internal/handler/userhandler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Use(middleware.LoggerData)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
	})

	router.Post("/user", h.CreateUser)

	router.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(env.Env.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Patch("/user{id}", h.UpdateUser)
		r.Get("/user/{id}", h.GetUserByID)
		r.Delete("/user/{id}", h.DeleteUser)
		r.Get("/user/list-all", h.FindManyUsers)
		r.Patch("/user/{id}/password", h.UpdateUserPassword)
	})

}

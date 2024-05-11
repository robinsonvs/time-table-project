package userhandler

import (
	"github.com/robinsonvs/time-table-project/internal/service/userservice"
	"net/http"
)

func NewUserHandler(service userservice.UserService) UserHandler {
	return &handler{
		service,
	}
}

type handler struct {
	service userservice.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindManyUsers(w http.ResponseWriter, r *http.Request)
	UpdateUserPassword(w http.ResponseWriter, r *http.Request)
}

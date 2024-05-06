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
}

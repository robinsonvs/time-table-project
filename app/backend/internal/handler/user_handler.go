package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/common/utils"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/handler/httperr"
	"github.com/robinsonvs/time-table-project/internal/handler/validation"
	"log/slog"
	"net/http"
)

// Create user
//
//	@Summary		Create new user
//	@Description	Endpoint for create user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateUserDto	true	"Create user dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/users [post]
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_user"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}

	err = h.userService.CreateUser(r.Context(), req)
	if err != nil {
		slog.Error(fmt.Sprintf("error to create user: %v", err), slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to create user")
		json.NewEncoder(w).Encode(msg)
		return
	}
}

// Update user
//
//	@Summary		Update user
//	@Description	Endpoint for update user
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.UpdateUserDto	false	"Update user dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/users [patch]
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateUserDto

	user, err := utils.DecodeJwt(r)
	if err != nil {
		slog.Error("error to decode jwt", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode jwt")
		json.NewEncoder(w).Encode(msg)
		return
	}
	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}
	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "userhandler"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}
	err = h.userService.UpdateUser(r.Context(), req, user.UUID)
	if err != nil {
		slog.Error(fmt.Sprintf("error to update user: %v", err), slog.String("package", "userhandler"))
		if err.Error() == "user not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("user not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		if err.Error() == "cep not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("cep not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		if err.Error() == "user already exists" {
			w.WriteHeader(http.StatusBadRequest)
			msg := httperr.NewBadRequestError("user already exists with this email")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
}

// User details
//
//	@Summary		User details
//	@Description	Get user by uuid
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"user uuid"
//	@Success		200	{object}	response.UserResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/users/{uuid} [get]
func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	uuidUser := chi.URLParam(r, "uuid")
	if uuidUser == "" {
		slog.Error("id is empty", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuidUserParser, err := uuid.Parse(uuidUser)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}

	res, err := h.userService.GetUserByID(r.Context(), uuidUserParser)
	if err != nil {
		slog.Error(fmt.Sprintf("error to get user: %v", err), slog.String("package", "userhandler"))
		if err.Error() == "user not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("user not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to get user")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Delete user
//
//	@Summary		Delete user
//	@Description	delete user by uuid
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"user uuid"
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/users/{uuid} [delete]
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = h.userService.DeleteUser(r.Context(), uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to delete user: %v", err), slog.String("package", "handler_user"))
		if err.Error() == "user not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("user not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to delete user")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Get many user
//
//	@Summary		Get many users
//	@Description	Get many users
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.ManyUsersResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/users/list-all [get]
func (h *handler) FindManyUsers(w http.ResponseWriter, r *http.Request) {
	res, err := h.userService.FindManyUsers(r.Context())
	if err != nil {
		slog.Error(fmt.Sprintf("error to find many users: %v", err), slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to find many users")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Update user password
//
//	@Summary		Update user password
//	@Description	Endpoint for Update user password
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid		path	string						true	"user uuid"
//	@Param			body	body	dto.UpdateUserPasswordDto	true	"Update user password dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/users/{uuid}/password [patch]
func (h *handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateUserPasswordDto

	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}
	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}
	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_user"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}
	err = h.userService.UpdateUserPassword(r.Context(), &req, uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to update user password: %v", err), slog.String("package", "handler_user"))
		if err.Error() == "user not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("user not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to update user password")
		json.NewEncoder(w).Encode(msg)
		return
	}

}

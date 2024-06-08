package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/handler/httperr"
	"github.com/robinsonvs/time-table-project/internal/handler/validation"
	"log/slog"
	"net/http"
	"strconv"
)

// Create discipline
//
//	@Summary		Create new discipline
//	@Description	Endpoint for create discipline
//	@Tags			discipline
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateDisciplineDto	true	"Create discipline dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/disciplines [post]
func (h *handler) CreateDiscipline(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateDisciplineDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_discipline"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}

	err = h.disciplineService.CreateDiscipline(r.Context(), req)
	if err != nil {
		if err.Error() == "discipline not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("discipline not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		slog.Error(fmt.Sprintf("error to create discipline: %v", err), slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

// Update discipline
//
//	@Summary		Update discipline
//	@Description	Endpoint for update discipline
//	@Tags			discipline
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid		path	string					true	"discipline uuid"
//	@Param			body	body	dto.UpdateDisciplineDto	false	"Update discipline dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/disciplines/{uuid} [patch]
func (h *handler) UpdateDiscipline(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateDisciplineDto

	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("discipline id is required", slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("discipline id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse discipline id: %v", err), slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("invalid discipline id")
		json.NewEncoder(w).Encode(msg)
		return
	}
	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}
	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_discipline"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}
	err = h.disciplineService.UpdateDiscipline(r.Context(), req, uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to update discipline: %v", err), slog.String("package", "handler_discipline"))
		if err.Error() == "discipline not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("discipline not found")
			json.NewEncoder(w).Encode(msg)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
}

// Discipline details
//
//	@Summary		Discipline details
//	@Description	Get discipline by uuid
//	@Tags			discipline
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"discipline uuid"
//	@Success		200	{object}	response.DisciplineResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/disciplines/{uuid} [get]
func (h *handler) GetDisciplineByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}

	res, err := h.disciplineService.GetDisciplineByID(r.Context(), uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to get discipline: %v", err), slog.String("package", "handler_discipline"))
		if err.Error() == "discipline not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("discipline not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to get discipline")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Delete discipline
//
//	@Summary		Delete discipline
//	@Description	delete discipline by uuid
//	@Tags			discipline
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"discipline uuid"
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/disciplines/{uuid} [delete]
func (h *handler) DeleteDiscipline(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = h.disciplineService.DeleteDiscipline(r.Context(), uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to delete discipline: %v", err), slog.String("package", "handler_discipline"))
		if err.Error() == "discipline not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("discipline not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to delete discipline")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Get many disciplines
//
//	@Summary		Get many disciplines
//	@Description	Get many disciplines
//	@Tags			discipline
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.ManyDisciplinesResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/disciplines/list-all [get]
func (h *handler) FindManyDisciplines(w http.ResponseWriter, r *http.Request) {
	res, err := h.disciplineService.FindManyDisciplines(r.Context())
	if err != nil {
		slog.Error(fmt.Sprintf("error to find many disciplines: %v", err), slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to find many disciplines")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Get many disciplines by course
//
//	@Summary		Get many disciplines by course
//	@Description	Get many disciplines by course
//	@Tags			discipline by course
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.ManyDisciplinesResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/disciplines/list-all/{courseId} [get]
func (h *handler) FindManyDisciplinesByCourseId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "courseId")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	courseId, _ := strconv.ParseInt(id, 10, 64)
	res, err := h.disciplineService.FindManyDisciplinesByCourseId(r.Context(), courseId)
	if err != nil {
		slog.Error(fmt.Sprintf("error to find many disciplines: %v", err), slog.String("package", "handler_discipline"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to find many disciplines")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

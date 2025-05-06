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

// Create availability
//
//	@Summary		Create new availability
//	@Description	Endpoint for create availability
//	@Tags			availability
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateAvailabilityDto	true	"Create availability dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/availabilities [post]
func (h *handler) CreateAvailability(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAvailabilityDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_availability"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}

	err = h.availabilityService.CreateAvailability(r.Context(), req)
	if err != nil {
		if err.Error() == "availability not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("availability not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		slog.Error(fmt.Sprintf("error to create availability: %v", err), slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

// Update availability
//
//	@Summary		Update availability
//	@Description	Endpoint for update availability
//	@Tags			availability
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid		path	string					true	"availability uuid"
//	@Param			body	body	dto.UpdateAvailabilityDto	false	"Update availability dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/availabilities/{uuid} [patch]
func (h *handler) UpdateAvailability(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateAvailabilityDto

	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("availability id is required", slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("availability id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse availability id: %v", err), slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("invalid availability id")
		json.NewEncoder(w).Encode(msg)
		return
	}
	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}
	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_availability"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}
	err = h.availabilityService.UpdateAvailability(r.Context(), req, uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to update availability: %v", err), slog.String("package", "handler_availability"))
		if err.Error() == "availability not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("availability not found")
			json.NewEncoder(w).Encode(msg)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
}

// Availability details
//
//	@Summary		Availability details
//	@Description	Get availability by uuid
//	@Tags			availability
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"availability uuid"
//	@Success		200	{object}	response.AvailabilityResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/availabilities/{uuid} [get]
func (h *handler) GetAvailabilityByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}

	res, err := h.availabilityService.GetAvailabilityByID(r.Context(), uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to get discipline: %v", err), slog.String("package", "handler_availability"))
		if err.Error() == "availability not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("availability not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to get availability")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Delete availability
//
//	@Summary		Delete availability
//	@Description	delete availability by uuid
//	@Tags			availability
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"availability uuid"
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/availabilities/{uuid} [delete]
func (h *handler) DeleteAvailability(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = h.availabilityService.DeleteAvailability(r.Context(), uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to delete availability: %v", err), slog.String("package", "handler_availability"))
		if err.Error() == "availability not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("availability not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to delete availability")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Get many availability
//
//	@Summary		Get many availability
//	@Description	Get many availability
//	@Tags			availability
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.ManyAvailabilitiesResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/availabilities/list-all [get]
func (h *handler) FindManyAvailabilities(w http.ResponseWriter, r *http.Request) {
	res, err := h.availabilityService.FindManyAvailabilities(r.Context())
	if err != nil {
		slog.Error(fmt.Sprintf("error to find many availabilities: %v", err), slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to find many availabilities")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Get many availabilities by professor
//
//	@Summary		Get many availabilities by professor
//	@Description	Get many availabilities by professor
//	@Tags			availabilities by professor
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			professorId	path	string	true	"availability professorId"
//	@Success		200	{object}	response.ManyAvailabilitiesResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/availabilities/list-all/{professorId} [get]
func (h *handler) FindManyAvailabilitiesByProfessorId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "professorId")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	professorId, _ := strconv.ParseInt(id, 10, 64)
	res, err := h.availabilityService.FindManyAvailabilitiesByProfessorId(r.Context(), professorId)
	if err != nil {
		slog.Error(fmt.Sprintf("error to find many availabilities: %v", err), slog.String("package", "handler_availability"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to find many availabilities")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

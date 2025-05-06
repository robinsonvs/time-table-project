package handler

import (
	"encoding/json"
	"fmt"
	"github.com/robinsonvs/time-table-project/internal/dto"
	"github.com/robinsonvs/time-table-project/internal/handler/httperr"
	"github.com/robinsonvs/time-table-project/internal/handler/validation"
	"log/slog"
	"net/http"
)

// Create eligible discipline
//
//	@Summary		Create new eligible discipline
//	@Description	Endpoint for create eligible discipline
//	@Tags			eligible discipline
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateEligibleDisciplineDto	true	"Create eligible discipline dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/eligible-disciplines [post]
func (h *handler) CreateEligibleDiscipline(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateEligibleDisciplineDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "handler_eligible_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_eligible_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_eligible_discipline"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}

	err = h.eligibleDisciplineService.CreateEligibleDiscipline(r.Context(), req)
	if err != nil {
		if err.Error() == "eligible discipline not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("eligible discipline not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		slog.Error(fmt.Sprintf("error to create eligible discipline: %v", err), slog.String("package", "handler_eligible_discipline"))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

// Delete eligible discipline
//
//	@Summary		Delete eligible discipline
//	@Description	delete eligible discipline by uuid
//	@Tags			eligible discipline
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.DeleteEligibleDisciplineDto	true	"Delete eligible discipline dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/eligible-disciplines [delete]
func (h *handler) DeleteEligibleDiscipline(w http.ResponseWriter, r *http.Request) {
	var req dto.DeleteEligibleDisciplineDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "handler_eligible_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_eligible_discipline"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_eligible_discipline"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}

	err = h.eligibleDisciplineService.DeleteEligibleDiscipline(r.Context(), req)
	if err != nil {
		slog.Error(fmt.Sprintf("error to delete eligible discipline: %v", err), slog.String("package", "handler_eligible_discipline"))
		if err.Error() == "eligible discipline not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("eligible discipline not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to delete eligible discipline")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

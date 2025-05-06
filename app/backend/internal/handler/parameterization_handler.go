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

// Create parameterization
//
//	@Summary		Create new parameterization
//	@Description	Endpoint for create parameterization
//	@Tags			parameterization
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateParameterizationDto	true	"Create parameterization dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/parameterizations [post]
func (h *handler) CreateParameterization(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateParameterizationDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_parameterization"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}

	err = h.parameterizationService.CreateParameterization(r.Context(), req)
	if err != nil {
		if err.Error() == "parameterization not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("parameterization not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		slog.Error(fmt.Sprintf("error to create parameterization: %v", err), slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

// Update parameterization
//
//	@Summary		Update parameterization
//	@Description	Endpoint for update parameterization
//	@Tags			parameterization
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid		path	string					true	"parameterization uuid"
//	@Param			body	body	dto.UpdateParameterizationDto	false	"Update parameterization dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/parameterizations/{uuid} [patch]
func (h *handler) UpdateParameterization(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateParameterizationDto

	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("parameterization id is required", slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("parameterization id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse parameterization id: %v", err), slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("invalid parameterization id")
		json.NewEncoder(w).Encode(msg)
		return
	}
	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}
	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_parameterization"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}
	err = h.parameterizationService.UpdateParameterization(r.Context(), req, uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to update parameterization: %v", err), slog.String("package", "handler_parameterization"))
		if err.Error() == "parameterization not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("parameterization not found")
			json.NewEncoder(w).Encode(msg)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
}

// Parameterization details
//
//	@Summary		Parameterization details
//	@Description	Get parameterization by uuid
//	@Tags			parameterization
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"parameterization uuid"
//	@Success		200	{object}	response.ParameterizationResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/parameterizations/{uuid} [get]
func (h *handler) GetParameterizationByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}

	res, err := h.parameterizationService.GetParameterizationByID(r.Context(), uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to get discipline: %v", err), slog.String("package", "handler_parameterization"))
		if err.Error() == "parameterization not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("parameterization not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to get parameterization")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Delete parameterization
//
//	@Summary		Delete parameterization
//	@Description	delete parameterization by uuid
//	@Tags			parameterization
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"parameterization uuid"
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/parameterizations/{uuid} [delete]
func (h *handler) DeleteParameterization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = h.parameterizationService.DeleteParameterization(r.Context(), uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to delete parameterization: %v", err), slog.String("package", "handler_parameterization"))
		if err.Error() == "parameterization not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("parameterization not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to delete parameterization")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Get many parameterization
//
//	@Summary		Get many parameterization
//	@Description	Get many parameterization
//	@Tags			parameterization
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.ManyParameterizationsResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/parameterizations/list-all [get]
func (h *handler) FindManyParameterizations(w http.ResponseWriter, r *http.Request) {
	res, err := h.parameterizationService.FindManyParameterizations(r.Context())
	if err != nil {
		slog.Error(fmt.Sprintf("error to find many parameterizations: %v", err), slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to find many parameterizations")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Get many parameterizations by semester
//
//	@Summary		Get many parameterizations by semester
//	@Description	Get many parameterizations by semester
//	@Tags			parameterizations by semester
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			semesterId	path	string	true	"parameterization semesterId"
//	@Success		200	{object}	response.ManyParameterizationsResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/parameterizations/list-all/{semesterId} [get]
func (h *handler) FindManyParameterizationsBySemesterId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "semesterId")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	semesterId, _ := strconv.ParseInt(id, 10, 64)
	res, err := h.parameterizationService.FindManyParameterizationsBySemesterId(r.Context(), semesterId)
	if err != nil {
		slog.Error(fmt.Sprintf("error to find many parameterizations: %v", err), slog.String("package", "handler_parameterization"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to find many parameterizations")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

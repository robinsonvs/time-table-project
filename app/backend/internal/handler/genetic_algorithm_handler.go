package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/robinsonvs/time-table-project/internal/handler/httperr"
	"log/slog"
	"net/http"
)

// Generate a proposal
//
//	@Summary		Generate new proposal
//	@Description	Endpoint for generate proposal
//	@Tags			proposal
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			parameterizationID	path	string	true	"parameterization uuid"
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/generate-proposal/{parameterizationID} [post]
func (h *handler) GenerateProposal(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "parameterizationID")
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

	err = h.geneticAlgorithmService.GenerateProposal(r.Context(), uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to generate a proposal: %v", err), slog.String("package", "handler_genetic"))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

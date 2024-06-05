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
)

// Create course
//
//	@Summary		Create new course
//	@Description	Endpoint for create course
//	@Tags			course
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateCourseDto	true	"Create course dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/courses [post]
func (h *handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCourseDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_course"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}

	err = h.courseService.CreateCourse(r.Context(), req)
	if err != nil {
		if err.Error() == "course not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("course not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		slog.Error(fmt.Sprintf("error to create course: %v", err), slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

// Update course
//
//	@Summary		Update course
//	@Description	Endpoint for update course
//	@Tags			course
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid		path	string					true	"course uuid"
//	@Param			body	body	dto.UpdateCourseDto	false	"Update course dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/courses/{uuid} [patch]
func (h *handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCourseDto

	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("course id is required", slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("course id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse course id: %v", err), slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("invalid course id")
		json.NewEncoder(w).Encode(msg)
		return
	}
	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}
	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_course"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}
	err = h.courseService.UpdateCourse(r.Context(), req, uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to update course: %v", err), slog.String("package", "handler_course"))
		if err.Error() == "course not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("course not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		//if err.Error() == "cep not found" {
		//	w.WriteHeader(http.StatusNotFound)
		//	msg := httperr.NewNotFoundError("cep not found")
		//	json.NewEncoder(w).Encode(msg)
		//	return
		//}
		//if err.Error() == "course already exists" {
		//	w.WriteHeader(http.StatusBadRequest)
		//	msg := httperr.NewBadRequestError("course already exists with this email")
		//	json.NewEncoder(w).Encode(msg)
		//	return
		//}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
}

// Course details
//
//	@Summary		Course details
//	@Description	Get course by uuid
//	@Tags			course
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"course uuid"
//	@Success		200	{object}	response.UserResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/courses/{uuid} [get]
func (h *handler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}

	res, err := h.courseService.GetCourseByID(r.Context(), uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to get course: %v", err), slog.String("package", "handler_course"))
		if err.Error() == "course not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("course not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to get course")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Delete course
//
//	@Summary		Delete course
//	@Description	delete course by uuid
//	@Tags			course
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"course uuid"
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/courses/{uuid} [delete]
func (h *handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = h.courseService.DeleteCourse(r.Context(), uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("error to delete course: %v", err), slog.String("package", "handler_course"))
		if err.Error() == "course not found" || err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("course not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to delete course")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Get many course
//
//	@Summary		Get many courses
//	@Description	Get many courses
//	@Tags			course
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.ManyCoursesResponse
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		404	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/courses/list-all [get]
func (h *handler) FindManyCourses(w http.ResponseWriter, r *http.Request) {
	res, err := h.courseService.FindManyCourses(r.Context())
	if err != nil {
		slog.Error(fmt.Sprintf("error to find many courses: %v", err), slog.String("package", "handler_course"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to find many courses")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type (
	AppointmentController interface {
		Add(ctx *gin.Context)
		GetAllAppointmentsByProfessionalId(ctx *gin.Context)
		GetAppointmentById(ctx *gin.Context)
	}

	AppointmentHandler struct {
		BaseHandler
		uc domains.AppointmentUseCase
	}

	request struct {
		ProfessionalID string    `json:"professional_id" validate:"required"`
		ServiceID      string    `json:"service_id" validate:"required"`
		ScheduledDate  time.Time `json:"schedule_date" validate:"required"`
		Email          string    `json:"user_email" validate:"required"`
		Phone          string    `json:"user_phone" validate:"required"`
		Notes          string    `json:"notes" validate:"required"`
		CreatedAt      time.Time `json:"created_at" validate:"required"`
		UpdatedAt      time.Time `json:"updated_at" validate:"required"`
	}
)

func NewAppointmentController(uc domains.AppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{uc: uc}
}

// Add godoc
// @Summary      Create an Appointment
// @Description  Create a new Appointment
// @Tags         Appointment
// @Accept       json
// @Produce      json
// @Param        establishment  body      domains.Establishment  true  "Appointment data"
// @Success      201            {object}  domains.HttpResponse{data=domains.Appointment}  "Appointment created successfully"
// @Failure      400            {object}  domains.HttpResponse  "Bad Request"
// @Failure      500            {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /appointments [post]
func (h *AppointmentHandler) Add(ctx *gin.Context) {
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	appointment := domains.Appointment{
		ProfessionalID: req.ProfessionalID,
		ServiceID:      req.ServiceID,
		ScheduledDate:  req.ScheduledDate,
		Email:          req.Email,
		Phone:          req.Phone,
		Notes:          req.Notes,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
	}

	appointmentCreated, err := h.uc.Add(ctx.Request.Context(), &appointment)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "Appointment created successfully", appointmentCreated)
}

// Add godoc
// @Summary      Get all appointments by professional id
// @Description  Get all appointments by professional ids
// @Tags         Appointment
// @Accept       json
// @Produce      json
// @Param        establishment  body      domains.Establishment  true  "Appointment data"
// @Success      201            {object}  domains.HttpResponse{data=domains.Appointment}  "appointment by id successfully retrieved"
// @Failure      400            {object}  domains.HttpResponse  "Bad Request"
// @Failure      500            {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /appointments/:id/professional [get]
func (h *AppointmentHandler) GetAllAppointmentsByProfessionalId(ctx *gin.Context) {
	id := ctx.Param("id")
	appointments, err := h.uc.GetAllAppointmentsByProfessionalId(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "appointments by professional id successfully retrieved", appointments)
}

// Add godoc
// @Summary      GetAppointmentById
// @Description  GetAppointmentById
// @Tags         Appointment
// @Accept       json
// @Produce      json
// @Param        establishment  body      domains.Establishment  true  "Appointment data"
// @Success      201            {object}  domains.HttpResponse{data=domains.Appointment}  "appointment by id successfully retrieved"
// @Failure      400            {object}  domains.HttpResponse  "Bad Request"
// @Failure      500            {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /appointments/:id [get]
func (h *AppointmentHandler) GetAppointmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	appointment, err := h.uc.GetAppointmentById(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "appointment by id successfully retrieved", appointment)
}

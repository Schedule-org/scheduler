package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type AppointmentController interface {
	Add(ctx *gin.Context)
	GetAllAppointmentsByProfessionalId(ctx *gin.Context)
	GetAppointmentById(ctx *gin.Context)
}
type AppointmentHandler struct {
	BaseHandler
	uc domains.AppointmentUseCase
}

func NewAppointmentController(uc domains.AppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{uc: uc}
}

func (h *AppointmentHandler) Add(ctx *gin.Context) {
	var input domains.Appointment
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	appointment, err := h.uc.Add(ctx.Request.Context(), &input)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "Appointment created successfully", appointment)
}

func (h *AppointmentHandler) GetAllAppointmentsByProfessionalId(ctx *gin.Context) {
	id := ctx.Param("id")
	appointments, err := h.uc.GetAllAppointmentsByProfessionalId(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "appointments by professional id successfully retrieved", appointments)
}

func (h *AppointmentHandler) GetAppointmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	appointment, err := h.uc.GetAppointmentById(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "appointment by id successfully retrieved", appointment)
}

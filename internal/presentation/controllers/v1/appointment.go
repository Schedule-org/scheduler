package controllers

import (
	"net/http"

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

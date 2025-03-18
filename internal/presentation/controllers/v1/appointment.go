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
	uc domains.AppointmentUseCase
}

func NewAppointmentController(uc domains.AppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{uc: uc}
}

func (h *AppointmentHandler) Add(ctx *gin.Context) {
	var input domains.Appointment
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.respondWithError(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	appointment, err := h.uc.Add(ctx.Request.Context(), &input)
	if err != nil {
		h.respondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.respondWithSuccess(ctx, http.StatusCreated, "Appointment created successfully", appointment)
}

func (h *AppointmentHandler) GetAllAppointmentsByProfessionalId(ctx *gin.Context) {
	id := ctx.Param("id")
	appointments, err := h.uc.GetAllAppointmentsByProfessionalId(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
		return
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "appointments by professional id successfully retrieved",
		Code:    http.StatusCreated,
		Data:    appointments,
	})
}

func (h *AppointmentHandler) GetAppointmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	appointment, err := h.uc.GetAppointmentById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
		return
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "appointment by id successfully retrieved",
		Code:    http.StatusCreated,
		Data:    appointment,
	})
}

func (h *AppointmentHandler) respondWithError(ctx *gin.Context, code int, message string, err error) {
	ctx.JSON(code, domains.HttpResponse{
		Message: message,
		Code:    code,
	})
}

func (h *AppointmentHandler) respondWithSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, domains.HttpResponse{
		Message: message,
		Code:    code,
		Data:    data,
	})
}

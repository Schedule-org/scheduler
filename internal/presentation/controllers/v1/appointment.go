package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domains"
)

type AppointmentController interface {
	Add(ctx *gin.Context)
}

type AppointmentUseCase struct {
	uc usecases.AppointmentUseCase
}

func NewAppointmentController(uc usecases.AppointmentUseCase) *AppointmentUseCase {
	return &AppointmentUseCase{uc: uc}
}

func (ctrl *AppointmentUseCase) Add(ctx *gin.Context) {
	var input domains.Appointment
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, domains.HttpResponse{
			Message: err.Error(),
		})
		return
	}

	output, err := ctrl.uc.Add(ctx.Request.Context(), &input)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
		return
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "appointment created successfully",
		Code:    http.StatusCreated,
		Data:    output,
	})
}

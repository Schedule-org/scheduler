package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type ProfessionalAvailabilityController interface {
	Add(ctx *gin.Context)
	GetProfessionalAvailabilityById(ctx *gin.Context)
}

type ProfessionalAvailabilityUseCase struct {
	uc domains.ProfessionalsAvailabilityUseCase
}

func NewProfessionalAvailabilityController(uc domains.ProfessionalsAvailabilityUseCase) *ProfessionalAvailabilityUseCase {
	return &ProfessionalAvailabilityUseCase{uc: uc}
}

func (ctrl *ProfessionalAvailabilityUseCase) Add(ctx *gin.Context) {
	var input domains.ProfessionalAvailability
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, domains.HttpResponse{
			Message: err.Error(),
		})
		return
	}

	availability, err := ctrl.uc.Add(ctx.Request.Context(), &input)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
		return
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "professional availability created successfully",
		Code:    http.StatusCreated,
		Data:    availability,
	})
}

func (ctrl *ProfessionalAvailabilityUseCase) GetProfessionalAvailabilityById(ctx *gin.Context) {
	professional_id := ctx.Param("id")
	availability, err := ctrl.uc.GetProfessionalAvailabilityById(ctx.Request.Context(), professional_id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
		return
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "professional availability retrieved successfully",
		Code:    http.StatusCreated,
		Data:    availability,
	})
}

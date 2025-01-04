package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domains"
)

type ServicesController interface {
	Add(ctx *gin.Context)
	FindServiceById(ctx *gin.Context)
	GetAllServicesByProfessionalId(ctx *gin.Context)
}

type ServicesUseCase struct {
	uc usecases.ServicesUseCase
}

func NewServicesController(uc usecases.ServicesUseCase) *ServicesUseCase {
	return &ServicesUseCase{uc: uc}
}

// AddService godoc
// @Summary      Add a new service
// @Description  Create a new service with the provided data
// @Tags         Services
// @Accept       json
// @Produce      json
// @Param        service  body      domains.Services  true  "Service data"
// @Success      201      {object}  domains.HttpResponse{data=domains.Services}  "Service created successfully"
// @Failure      400      {object}  domains.HttpResponse  "Bad Request"
// @Failure      500      {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /services [post]
func (ctrl *ServicesUseCase) Add(ctx *gin.Context) {
	var input domains.Services
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
		Message: "service created successfully",
		Code:    http.StatusCreated,
		Data:    output,
	})
}

// FindServiceById godoc
// @Summary      Find service by ID
// @Description  Retrieve a service using its unique ID
// @Tags         Services
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Service ID"
// @Success      200  {object}  domains.HttpResponse{data=domains.Services}  "Service found successfully"
// @Failure      404  {object}  domains.HttpResponse  "Service not found"
// @Failure      500  {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /service_id/{id} [get]
func (ctrl *ServicesUseCase) FindServiceById(ctx *gin.Context) {
	id := ctx.Param("id")
	output, err := ctrl.uc.FindServiceById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "service found successfully",
		Code:    http.StatusOK,
		Data:    output,
	})
}

func (ctrl *ServicesUseCase) GetAllServicesByProfessionalId(ctx *gin.Context) {
	professional_id := ctx.Param("profissional_id")
	output, err := ctrl.uc.GetAllServicesByProfessionalId(ctx.Request.Context(), professional_id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "services found successfully",
		Code:    http.StatusOK,
		Data:    output,
	})
}

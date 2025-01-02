package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domains"
)

type EstablishmentController interface {
	Add(ctx *gin.Context)
	FindEstablishmentById(ctx *gin.Context)
}

type EstablishmentUseCase struct {
	uc usecases.EstablishmentUseCase
}

func NewEstablishmentController(uc usecases.EstablishmentUseCase) *EstablishmentUseCase {
	return &EstablishmentUseCase{uc: uc}
}

// Add godoc
// @Summary      Add a new establishment
// @Description  Create a new establishment with the provided data
// @Tags         Establishments
// @Accept       json
// @Produce      json
// @Param        establishment  body      domains.Establishment  true  "Establishment data"
// @Success      201            {object}  domains.HttpResponse{data=domains.Establishment}  "Establishment created successfully"
// @Failure      400            {object}  domains.HttpResponse  "Bad Request"
// @Failure      500            {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /establishments [post]
func (ctrl *EstablishmentUseCase) Add(ctx *gin.Context) {
	var input domains.Establishment
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
		Message: "Establishment created successfully",
		Code:    http.StatusCreated,
		Data:    output,
	})
}

// FindEstablishmentById godoc
// @Summary      Find establishment by ID
// @Description  Retrieve an establishment using its unique ID
// @Tags         Establishments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Establishment ID"
// @Success      200  {object}  domains.HttpResponse{data=domains.Establishment}  "Establishment found successfully"
// @Failure      404  {object}  domains.HttpResponse  "Establishment not found"
// @Failure      500  {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /establishments/{id} [get]
func (ctrl *EstablishmentUseCase) FindEstablishmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	output, err := ctrl.uc.FindEstablishmentById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "Establishment found successfully",
		Code:    http.StatusOK,
		Data:    output,
	})
}

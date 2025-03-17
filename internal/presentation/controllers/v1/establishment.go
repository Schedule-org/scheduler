package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type EstablishmentController interface {
	Add(ctx *gin.Context)
	FindEstablishmentById(ctx *gin.Context)
	GetAllProfessinalsByEstablishmentId(ctx *gin.Context)
	UpdateEstablishmentById(ctx *gin.Context)
	GetEstablishmentReport(ctx *gin.Context)
}

type EstablishmentUseCase struct {
	uc domains.EstablishmentUseCase
}

func NewEstablishmentController(uc domains.EstablishmentUseCase) *EstablishmentUseCase {
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

	establishment, err := ctrl.uc.Add(ctx.Request.Context(), &input)
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
		Data:    establishment,
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
// @Router       /establishment_id/{id} [get]
func (ctrl *EstablishmentUseCase) FindEstablishmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	establishment, err := ctrl.uc.FindEstablishmentById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "Establishment found successfully",
		Code:    http.StatusOK,
		Data:    establishment,
	})
}

// GetAllProfessinalsByEstablishmentId godoc
// @Summary      Get All professionals by establishment ID
// @Description  Retrieve professionals using unique ID
// @Tags         Establishments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "id"
// @Success      200  {object}  domains.HttpResponse{data=domains.Establishment}  "Establishment found successfully"
// @Failure      404  {object}  domains.HttpResponse  "Professionals found successfully"
// @Failure      500  {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /establishment/:id/professionals [get]
func (ctrl *EstablishmentUseCase) GetAllProfessinalsByEstablishmentId(ctx *gin.Context) {
	establishment_id := ctx.Param("id")
	professionals, err := ctrl.uc.GetAllProfessionalsByEstablishmentId(ctx.Request.Context(), establishment_id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "Professionals found successfully",
		Code:    http.StatusOK,
		Data:    professionals,
	})
}

func (ctrl *EstablishmentUseCase) UpdateEstablishmentById(ctx *gin.Context) {
	establishment_id := ctx.Param("id")
	var input domains.Establishment
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, domains.HttpResponse{
			Message: err.Error(),
		})
		return
	}
	establishments, err := ctrl.uc.UpdateEstablishmentById(ctx.Request.Context(), establishment_id, &input)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "Establishment update successfully",
		Code:    http.StatusOK,
		Data:    establishments,
	})
}

func (ctrl *EstablishmentUseCase) GetEstablishmentReport(ctx *gin.Context) {
	establishment_id := ctx.Param("id")
	establishmentReport, err := ctrl.uc.GetEstablishmentReport(ctx.Request.Context(), establishment_id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "Establishment report",
		Code:    http.StatusOK,
		Data:    establishmentReport,
	})
}

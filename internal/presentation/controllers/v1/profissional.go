package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domains"
)

type ProfessionalsController interface {
	Add(ctx *gin.Context)
	FindEstablishmentById(ctx *gin.Context)
}

type ProfessionalsUseCase struct {
	uc usecases.ProfessionalsUseCase
}

func NewProfessionalController(uc usecases.ProfessionalsUseCase) *ProfessionalsUseCase {
	return &ProfessionalsUseCase{uc: uc}
}

// Add godoc
// @Summary      Add a new professional
// @Description  Create a new professional with the provided data
// @Tags         Professionals
// @Accept       json
// @Produce      json
// @Param        professional  body      domains.Professionals  true  "Professional data"
// @Success      201           {object}  domains.HttpResponse{data=domains.Professionals}  "Professional created successfully"
// @Failure      400           {object}  domains.HttpResponse  "Bad Request"
// @Failure      500           {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /professionals [post]
func (ctrl *ProfessionalsUseCase) Add(ctx *gin.Context) {
	var input domains.Professionals
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
		Message: "Professional created successfully",
		Code:    http.StatusCreated,
		Data:    output,
	})
}

// FindProfessionalById godoc
// @Summary      Find professional by ID
// @Description  Retrieve a professional using its unique ID
// @Tags         Professionals
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Professional ID"
// @Success      200  {object}  domains.HttpResponse{data=domains.Professionals}  "Professional found successfully"
// @Failure      404  {object}  domains.HttpResponse  "Professional not found"
// @Failure      500  {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /professionals/{id} [get]
func (ctrl *ProfessionalsUseCase) FindEstablishmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	output, err := ctrl.uc.FindProfessionalById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "Professional found successfully",
		Code:    http.StatusOK,
		Data:    output,
	})
}

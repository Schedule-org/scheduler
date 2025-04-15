package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type (
	ProfessionalsHandler struct {
		BaseHandler
		uc domains.ProfessionalsUseCase
	}

	professionalRequest struct {
		Name            string `json:"name" validate:"required"`
		Role            string `json:"role" validate:"required"`
		EstablishmentId string `json:"establishment" validate:"required"`
	}
)

func NewProfessionalController(uc domains.ProfessionalsUseCase) *ProfessionalsHandler {
	return &ProfessionalsHandler{uc: uc}
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
func (h *ProfessionalsHandler) Add(ctx *gin.Context) {
	var req professionalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	professionalCreated := domains.Professionals{
		Name:            req.Name,
		Role:            req.Role,
		EstablishmentId: req.EstablishmentId,
	}
	professional, err := h.uc.Add(ctx.Request.Context(), &professionalCreated)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "professional created successfully", professional)
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
func (h *ProfessionalsHandler) FindProfessionalById(ctx *gin.Context) {
	id := ctx.Param("id")
	professional, err := h.uc.FindProfessionalById(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "professional found successfully", professional)
}

// UpdateProfessionalById godoc
// @Summary      UpdateProfessionalById
// @Description  Retrieve a professional using its unique ID
// @Tags         Professionals
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Professional ID"
// @Success      200  {object}  domains.HttpResponse{data=domains.Professionals}  "professional update successfully"
// @Failure      404  {object}  domains.HttpResponse  "Professional not found"
// @Failure      500  {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /professionals/{id} [get]
func (h *ProfessionalsHandler) UpdateProfessionalById(ctx *gin.Context) {
	var input domains.Professionals
	id := ctx.Param("id")
	professional, err := h.uc.UpdateProfessionalById(ctx.Request.Context(), id, &input)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "professional update successfully", professional)
}

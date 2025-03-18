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

type ProfessionalAvailabilityHandler struct {
	BaseHandler
	uc domains.ProfessionalsAvailabilityUseCase
}

func NewProfessionalAvailabilityController(uc domains.ProfessionalsAvailabilityUseCase) *ProfessionalAvailabilityHandler {
	return &ProfessionalAvailabilityHandler{uc: uc}
}

func (h *ProfessionalAvailabilityHandler) Add(ctx *gin.Context) {
	var input domains.ProfessionalAvailability
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	availability, err := h.uc.Add(ctx.Request.Context(), &input)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "professional availability created successfully", availability)
}

func (h *ProfessionalAvailabilityHandler) GetProfessionalAvailabilityById(ctx *gin.Context) {
	professional_id := ctx.Param("id")
	availability, err := h.uc.GetProfessionalAvailabilityById(ctx.Request.Context(), professional_id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "professional availability retrieved successfully", availability)
}

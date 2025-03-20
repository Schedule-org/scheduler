package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type (
	ProfessionalAvailabilityController interface {
		Add(ctx *gin.Context)
		GetProfessionalAvailabilityById(ctx *gin.Context)
	}

	ProfessionalAvailabilityHandler struct {
		BaseHandler
		uc domains.ProfessionalsAvailabilityUseCase
	}
)

func NewProfessionalAvailabilityController(uc domains.ProfessionalsAvailabilityUseCase) *ProfessionalAvailabilityHandler {
	return &ProfessionalAvailabilityHandler{uc: uc}
}

// Add godoc
// @Summary      Add
// @Description  Add
// @Tags         ProfessionalAvailability
// @Accept       json
// @Produce      json
// @Param        establishment  body      domains.Establishment  true  "ProfessionalAvailabilityHandler data"
// @Success      201            {object}  domains.HttpResponse{data=domains.ProfessionalAvailability}  "professional availability created successfully by id successfully retrieved"
// @Failure      400            {object}  domains.HttpResponse  "Bad Request"
// @Failure      500            {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /availability/ [post]
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

// Add godoc
// @Summary      Add
// @Description  Add
// @Tags         ProfessionalAvailability
// @Accept       json
// @Produce      json
// @Param        establishment  body      domains.Establishment  true  "ProfessionalAvailabilityHandler data"
// @Success      201            {object}  domains.HttpResponse{data=domains.ProfessionalAvailability}  "professional availability retrieved successfullys availability created successfully by id successfully retrieved"
// @Failure      400            {object}  domains.HttpResponse  "Bad Request"
// @Failure      500            {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /availability/:id/professional [get]
func (h *ProfessionalAvailabilityHandler) GetProfessionalAvailabilityById(ctx *gin.Context) {
	professional_id := ctx.Param("id")
	availability, err := h.uc.GetProfessionalAvailabilityById(ctx.Request.Context(), professional_id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "professional availability retrieved successfully", availability)
}

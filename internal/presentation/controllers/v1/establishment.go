package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type (
	EstablishmentController interface {
		Add(ctx *gin.Context)
		FindEstablishmentById(ctx *gin.Context)
		GetAllProfessinalsByEstablishmentId(ctx *gin.Context)
		UpdateEstablishmentById(ctx *gin.Context)
		GetEstablishmentReport(ctx *gin.Context)
	}

	EstablishmentHandler struct {
		BaseHandler
		uc domains.EstablishmentUseCase
	}
)

func NewEstablishmentController(uc domains.EstablishmentUseCase) *EstablishmentHandler {
	return &EstablishmentHandler{uc: uc}
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
func (h *EstablishmentHandler) Add(ctx *gin.Context) {
	var input domains.Establishment
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	establishment, err := h.uc.Add(ctx.Request.Context(), &input)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}
	h.RespondWithSuccess(ctx, http.StatusCreated, "Establishment created successfully", establishment)
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
func (h *EstablishmentHandler) FindEstablishmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.RespondWithError(ctx, http.StatusBadRequest, "id is required", nil)
		return
	}
	establishment, err := h.uc.FindEstablishmentById(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "Establishment found successfully", establishment)
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
func (h *EstablishmentHandler) GetAllProfessinalsByEstablishmentId(ctx *gin.Context) {
	establishment_id := ctx.Param("id")
	if establishment_id == "" {
		h.RespondWithError(ctx, http.StatusBadRequest, "establishment_id is required", nil)
		return
	}
	professionals, err := h.uc.GetAllProfessionalsByEstablishmentId(ctx.Request.Context(), establishment_id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}
	h.RespondWithSuccess(ctx, http.StatusOK, "Professionals found successfully", professionals)
}

func (h *EstablishmentHandler) UpdateEstablishmentById(ctx *gin.Context) {
	establishment_id := ctx.Param("id")
	if establishment_id == "" {
		h.RespondWithError(ctx, http.StatusBadRequest, "establishment_id is required", nil)
		return
	}
	var input domains.Establishment
	establishments, err := h.uc.UpdateEstablishmentById(ctx.Request.Context(), establishment_id, &input)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "Establishment update successfully", establishments)
}

func (h *EstablishmentHandler) GetEstablishmentReport(ctx *gin.Context) {
	establishment_id := ctx.Param("id")
	establishmentReport, err := h.uc.GetEstablishmentReport(ctx.Request.Context(), establishment_id)
	if err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "Establishment report", establishmentReport)
}

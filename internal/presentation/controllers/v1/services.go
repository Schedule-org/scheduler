package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type (
	ServicesController interface {
		Add(ctx *gin.Context)
		FindServiceById(ctx *gin.Context)
		GetAllServicesByProfessionalId(ctx *gin.Context)
	}

	ServicesHandler struct {
		BaseHandler
		uc domains.ServicesUseCase
	}
)

func NewServicesController(uc domains.ServicesUseCase) *ServicesHandler {
	return &ServicesHandler{uc: uc}
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
func (h *ServicesHandler) Add(ctx *gin.Context) {
	var input domains.Services
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	service, err := h.uc.Add(ctx.Request.Context(), &input)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "service created successfully", service)
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
func (h *ServicesHandler) FindServiceById(ctx *gin.Context) {
	id := ctx.Param("id")
	service, err := h.uc.FindServiceById(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "service found successfully", service)
}

func (h *ServicesHandler) GetAllServicesByProfessionalId(ctx *gin.Context) {
	professional_id := ctx.Param("id")
	services, err := h.uc.GetAllServicesByProfessionalId(ctx.Request.Context(), professional_id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "all services found successfully", services)
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	AccountHandler struct {
		BaseHandler
		uc domain.AccountUseCase
	}

	accountRequest struct {
		Cnpj  string `json:"cnpj" validate:"required"`
		Email string `json:"email" validate:"required"`
	}
)

func NewAccountController(uc domain.AccountUseCase) *AccountHandler {
	return &AccountHandler{uc: uc}
}

// Add godoc
// @Summary      Create an Account
// @Description  Create a new Account
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        establishment  body      domain.Establishment  true  "Account data"
// @Success      201            {object}  domain.HttpResponse{data=domain.Account}  "Account created successfully"
// @Failure      400            {object}  domain.HttpResponse  "Bad Request"
// @Failure      500            {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /accounts [post]
func (h *AccountHandler) Add(ctx *gin.Context) {
	var req accountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	appointment := domain.Account{
		Email: req.Email,
		Cnpj:  req.Cnpj,
	}

	account, err := h.uc.Add(ctx.Request.Context(), &appointment)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "Account created successfully", account)
}

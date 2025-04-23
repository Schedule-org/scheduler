package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	UserHandler struct {
		BaseHandler
		uc domain.UserUseCase
	}

	userRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)

func NewUserController(uc domain.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

// Add godoc
// @Summary      Add a new user
// @Description  Create a new user with the provided data
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      domain.User  true  "User data"
// @Success      201   {object}  domain.HttpResponse{data=dto.UserDTO}  "User created successfully"
// @Failure      400   {object}  domain.HttpResponse  "Bad Request"
// @Failure      500   {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /users [post]
func (h *UserHandler) Add(ctx *gin.Context) {
	var req userRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	userCreated := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := h.uc.Add(ctx.Request.Context(), &userCreated)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "User created successfully", user)
}

// FindUserById godoc
// @Summary      Find a user by ID
// @Description  Retrieve a user by their unique ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "User ID"
// @Success      200  {object}  domain.HttpResponse{data=dto.UserDTO}  "User found successfully"
// @Failure      400  {object}  domain.HttpResponse  "Bad Request"
// @Failure      404  {object}  domain.HttpResponse  "User not found"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /users/{id} [get]
func (h *UserHandler) FindUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.uc.FindUserById(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "User found successfully", user)
}

// FindAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve a list of all users in the system
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.HttpResponse{data=[]dto.UserDTO}  "Users retrieved successfully"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /users [get]
func (h *UserHandler) FindAllUsers(ctx *gin.Context) {
	users, err := h.uc.FindAllUsers(ctx.Request.Context())
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "Users retrieved", users)
}

func (h *UserHandler) FindAllEstablishmentsByUserId(ctx *gin.Context) {
	id := ctx.Param("id")
	establishments, err := h.uc.FindAllEstablishmentsByUserId(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "Users establishments retrieved", establishments)
}

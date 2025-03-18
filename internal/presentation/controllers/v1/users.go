package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type UserController interface {
	Add(ctx *gin.Context)
	FindUserById(ctx *gin.Context)
	FindAllUsers(ctx *gin.Context)
	FindAllEstablishmentsByUserId(ctx *gin.Context)
}

type UserHandler struct {
	uc domains.UserUseCase
}

func NewUserController(uc domains.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

// Add godoc
// @Summary      Add a new user
// @Description  Create a new user with the provided data
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      domains.User  true  "User data"
// @Success      201   {object}  domains.HttpResponse{data=dto.UserDTO}  "User created successfully"
// @Failure      400   {object}  domains.HttpResponse  "Bad Request"
// @Failure      500   {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /users [post]
func (h *UserHandler) Add(ctx *gin.Context) {
	var input domains.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, domains.HttpResponse{
			Message: err.Error(),
		})
		return
	}

	users, err := h.uc.Add(ctx.Request.Context(), &input)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
		return
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "User created successfully",
		Code:    http.StatusCreated,
		Data:    users,
	})
}

// FindUserById godoc
// @Summary      Find a user by ID
// @Description  Retrieve a user by their unique ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "User ID"
// @Success      200  {object}  domains.HttpResponse{data=dto.UserDTO}  "User found successfully"
// @Failure      400  {object}  domains.HttpResponse  "Bad Request"
// @Failure      404  {object}  domains.HttpResponse  "User not found"
// @Failure      500  {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /users/{id} [get]
func (h *UserHandler) FindUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	users, err := h.uc.FindUserById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "User found successfully",
		Code:    http.StatusOK,
		Data:    users,
	})
}

// FindAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve a list of all users in the system
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  domains.HttpResponse{data=[]dto.UserDTO}  "Users retrieved successfully"
// @Failure      500  {object}  domains.HttpResponse  "Internal Server Error"
// @Router       /users [get]
func (h *UserHandler) FindAllUsers(ctx *gin.Context) {
	users, err := h.uc.FindAllUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "Users retrieved",
		Code:    http.StatusOK,
		Data:    users,
	})
}

func (h *UserHandler) FindAllEstablishmentsByUserId(ctx *gin.Context) {
	id := ctx.Param("id")
	establishments, err := h.uc.FindAllEstablishmentsByUserId(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "Users establishments retrieved",
		Code:    http.StatusOK,
		Data:    establishments,
	})
}

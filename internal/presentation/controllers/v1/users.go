package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/dto"
)

type UserController interface {
	Add(ctx *gin.Context)
	FindUserById(ctx *gin.Context)
	FindAllUsers(ctx *gin.Context)
}

type UserUseCase struct {
	uc usecases.AddUserUseCase
}

func NewUserController(uc usecases.AddUserUseCase) *UserUseCase {
	return &UserUseCase{uc: uc}
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
func (ctrl *UserUseCase) Add(ctx *gin.Context) {
	var input domains.User
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

	user := dto.MapToUserDTO(output)
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "User created successfully",
		Code:    http.StatusCreated,
		Data:    user,
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
func (ctrl *UserUseCase) FindUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	output, err := ctrl.uc.FindUserById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}
	user := dto.MapToUserDTO(output)
	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "User found successfully",
		Data:    user,
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
func (ctrl *UserUseCase) FindAllUsers(ctx *gin.Context) {
	output, err := ctrl.uc.FindAllUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(err.Code, domains.HttpResponse{
			Message: err.Message,
			Code:    err.Code,
		})
	}

	var users []dto.UserDTO

	for i := 0; i < len(output); i++ {
		mapUser := dto.MapToUserDTO(&output[i])
		users = append(users, *mapUser)
	}

	ctx.JSON(http.StatusOK, domains.HttpResponse{
		Message: "Users retrieved",
		Code:    http.StatusOK,
		Data:    users,
	})
}

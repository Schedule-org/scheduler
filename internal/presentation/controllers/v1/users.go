package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/tadix-backend/internal/core/usecases"
	"github.com/hebertzin/tadix-backend/internal/domains"
	"github.com/hebertzin/tadix-backend/internal/infra/dto"
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

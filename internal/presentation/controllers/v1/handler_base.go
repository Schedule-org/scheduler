package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type BaseHandler struct{}

func (b *BaseHandler) RespondWithError(ctx *gin.Context, code int, message string, err error) {
	ctx.JSON(code, domains.HttpResponse{
		Message: message,
		Code:    code,
	})
}

func (b *BaseHandler) RespondWithSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, domains.HttpResponse{
		Message: message,
		Code:    code,
		Data:    data,
	})
}

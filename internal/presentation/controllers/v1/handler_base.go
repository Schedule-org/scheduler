package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

type BaseHandler struct{}

// response following RFC 9457 format
// This standardized structure provides a machine-readable and human-readable error format.
// More info: https://www.rfc-editor.org/rfc/rfc9457.html
func (b *BaseHandler) RespondWithError(ctx *gin.Context, status int, title string, err error) {
	ctx.JSON(status, domains.ErrorResponse{
		Title:    title,
		Status:   status,
		Instance: ctx.Request.URL.String(),
	})
}

func (b *BaseHandler) RespondWithSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, domains.HttpResponse{
		Message: message,
		Status:  code,
		Data:    data,
	})
}

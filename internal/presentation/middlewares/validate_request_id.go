package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domains"
)

func ValidateParamRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := strings.TrimSpace(ctx.Param("id"))
		if id == "" {
			ctx.JSON(http.StatusBadRequest, domains.HttpResponse{
				Message: "The 'id' parameter is required",
				Status:  http.StatusBadRequest,
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

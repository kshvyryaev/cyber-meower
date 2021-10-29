package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryHandlerMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(context *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Message: err})
			return
		}

		context.AbortWithStatus(http.StatusInternalServerError)
	})
}

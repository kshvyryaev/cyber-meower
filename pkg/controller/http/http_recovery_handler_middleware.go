package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HttpRecoveryHandlerMiddleware struct {
	logger *zap.Logger
}

func ProvideHttpRecoveryHandlerMiddleware(logger *zap.Logger) *HttpRecoveryHandlerMiddleware {
	return &HttpRecoveryHandlerMiddleware{
		logger: logger,
	}
}

func (handler *HttpRecoveryHandlerMiddleware) Handle() gin.HandlerFunc {
	return gin.CustomRecovery(func(context *gin.Context, recovered interface{}) {
		err, ok := recovered.(string)

		if ok {
			context.AbortWithStatusJSON(http.StatusInternalServerError, HttpErrorResponse{Message: err})
		} else {
			context.AbortWithStatus(http.StatusInternalServerError)
		}

		handler.logger.Error("panic happend: " + err)
	})
}

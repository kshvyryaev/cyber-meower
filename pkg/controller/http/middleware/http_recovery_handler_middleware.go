package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/http/response"
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
			context.AbortWithStatusJSON(http.StatusInternalServerError, response.HttpErrorResponse{Message: err})
		} else {
			context.AbortWithStatus(http.StatusInternalServerError)
		}

		handler.logger.Error("panic happend", zap.String("error", err))
	})
}

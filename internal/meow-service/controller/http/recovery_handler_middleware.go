package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type RecoveryHandlerMiddleware struct {
	logger *zap.Logger
}

func (handler *RecoveryHandlerMiddleware) Handle() gin.HandlerFunc {
	return gin.CustomRecovery(func(context *gin.Context, recovered interface{}) {
		err, ok := recovered.(string)

		if ok {
			context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Message: err})
		} else {
			context.AbortWithStatus(http.StatusInternalServerError)
		}

		handler.logger.Error("panic happend: " + err)
	})
}

func ProvideRecoveryHandlerMiddleware(logger *zap.Logger) *RecoveryHandlerMiddleware {
	return &RecoveryHandlerMiddleware{
		logger: logger,
	}
}

var RecoveryHandlerMiddlewareSet = wire.NewSet(
	ProvideRecoveryHandlerMiddleware,
)

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/http/response"
	"go.uber.org/zap"
)

type HttpErrorHandlerMiddleware struct {
	logger *zap.Logger
}

func ProvideHttpErrorHandlerMiddleware(logger *zap.Logger) *HttpErrorHandlerMiddleware {
	return &HttpErrorHandlerMiddleware{
		logger: logger,
	}
}

func (handler *HttpErrorHandlerMiddleware) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			err := context.Errors[0].Err
			context.AbortWithStatusJSON(http.StatusInternalServerError, response.HttpErrorResponse{Message: err.Error()})

			handler.logger.Error("error happend", zap.Error(err))
		}
	}
}

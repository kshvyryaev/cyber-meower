package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
			context.AbortWithStatusJSON(http.StatusInternalServerError, HttpErrorResponse{Message: err.Error()})

			handler.logger.Error("error happend: " + err.Error())
		}
	}
}

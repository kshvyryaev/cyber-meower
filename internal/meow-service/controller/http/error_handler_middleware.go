package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type ErrorHandlerMiddleware struct {
	Logger *zap.Logger
}

func (handler *ErrorHandlerMiddleware) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			err := context.Errors[0].Err
			context.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})

			handler.Logger.Error("error happend: " + err.Error())
		}
	}
}

func ProvideErrorHandlerMiddleware(logger *zap.Logger) *ErrorHandlerMiddleware {
	return &ErrorHandlerMiddleware{
		Logger: logger,
	}
}

var ErrorHandlerMiddlewareSet = wire.NewSet(
	ProvideErrorHandlerMiddleware,
)

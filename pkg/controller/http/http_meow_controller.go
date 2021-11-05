package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/command"
)

type HttpMeowController struct {
	createMeowCommandHandler *command.СreateMeowCommandHandler
}

func ProvideHttpMeowController(createMeowCommandHandler *command.СreateMeowCommandHandler) *HttpMeowController {
	return &HttpMeowController{
		createMeowCommandHandler: createMeowCommandHandler,
	}
}

func (controller *HttpMeowController) Route(router *gin.Engine) {
	meow := router.Group("/meow")
	{
		meow.POST("/", controller.Create)
	}
}

func (controller *HttpMeowController) Create(context *gin.Context) {
	var command command.CreateMeowCommand
	if err := context.BindJSON(&command); err != nil {
		context.Error(err)
		return
	}

	response, err := controller.createMeowCommandHandler.Handle(&command)
	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, response)
}

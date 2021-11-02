package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/command"
)

type MeowController struct {
	createMeowCommandHandler *command.СreateMeowCommandHandler
}

func ProvideMeowController(createMeowCommandHandler *command.СreateMeowCommandHandler) *MeowController {
	return &MeowController{
		createMeowCommandHandler: createMeowCommandHandler,
	}
}

func (controller *MeowController) Route(router *gin.Engine) {
	meow := router.Group("/meow")
	{
		meow.POST("/", controller.CreateMeow)
	}
}

func (controller *MeowController) CreateMeow(context *gin.Context) {
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

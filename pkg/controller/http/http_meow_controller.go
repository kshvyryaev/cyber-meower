package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/http/request"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/http/response"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/domain"
)

type HttpMeowController struct {
	usecase domain.MeowUsecase
}

func ProvideHttpMeowController(usecase domain.MeowUsecase) *HttpMeowController {
	return &HttpMeowController{
		usecase: usecase,
	}
}

func (controller *HttpMeowController) Route(router *gin.Engine) {
	meow := router.Group("/meow")
	{
		meow.POST("/", controller.Create)
	}
}

func (controller *HttpMeowController) Create(context *gin.Context) {
	var request request.CreateMeowRequest
	if err := context.BindJSON(&request); err != nil {
		context.Error(err)
		return
	}

	id, err := controller.usecase.Create(request.Body)
	if err != nil {
		context.Error(err)
		return
	}

	response := response.CreateMeowResponse{
		ID: id,
	}

	context.JSON(http.StatusOK, response)
}

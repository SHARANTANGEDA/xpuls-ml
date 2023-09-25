package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/xpuls-com/xpuls-ml/dto"
)

func ProjectRoutes(route *gin.RouterGroup) {

	route.POST("/create", tonic.Handler(dto.ProjectService.Create, 200))
}

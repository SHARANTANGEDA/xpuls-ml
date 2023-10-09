package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/xpuls-com/xpuls-ml/dto"
)

func ProjectRoutes(route *gin.RouterGroup) {

	route.POST("/create", tonic.Handler(dto.ProjectService.Create, 200))
	route.GET("/", tonic.Handler(dto.ProjectService.List, 200))
	route.GET("/is-slug-available",
		tonic.Handler(dto.ProjectService.CheckIfProjectSlugAvailable, 200))
	//route.GET("/:id", tonic.Handler(dto.ProjectService.GetById, 200))
}

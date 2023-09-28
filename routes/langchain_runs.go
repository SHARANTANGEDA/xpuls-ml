package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/xpuls-com/xpuls-ml/services"
)

func LangChainRuns(route *gin.RouterGroup) {

	route.POST("/runs", tonic.Handler(services.LangChainRunsServiceV2.TrackRun, 200))
	route.PATCH("/runs/:id", tonic.Handler(services.LangChainRunsServiceV2.PatchRun, 200))
	//route.GET("/", tonic.Handler(dto.ProjectService.List, 200))
	//route.GET("/:id", tonic.Handler(dto.ProjectService.GetById, 200))
}

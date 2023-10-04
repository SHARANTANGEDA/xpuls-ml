package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/xpuls-com/xpuls-ml/services"
)

func LangChainRuns(route *gin.RouterGroup) {
	route.GET("/:project_id/runs",
		tonic.Handler(services.LangChainRunsServiceV2.GetRunsInProject, 200))
	route.GET("/:project_id/runs/:chain_id",
		tonic.Handler(services.LangChainRunsStepsService.GetRunsByChainId, 200))
	route.GET("/:project_id/runs/filters/keys",
		tonic.Handler(services.LangChainRunsServiceV2.GetRunFilterKeys, 200))

	route.GET("/:project_id/runs/filters/values",
		tonic.Handler(services.LangChainRunsServiceV2.GetRunFilterKeyValues, 200))
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/xpuls-com/xpuls-ml/services"
)

func LangChainRunsTrack(route *gin.RouterGroup) {

	route.POST("/runs", tonic.Handler(services.LangChainRunsServiceV2.AddLangChainRunToQueue,
		200))
	route.PATCH("/runs/:id", tonic.Handler(services.LangChainRunsServiceV2.AddLangChainRunPatchToQueue,
		200))
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/xpuls-com/xpuls-ml/services"
)

func PromptRegistry(route *gin.RouterGroup) {
	route.GET("/:project_id/prompt/:prompt_id/versions",
		tonic.Handler(services.PromptRegistryService.ListPromptVersions, 200))
	route.GET("/:project_id/prompt/:prompt_id/latest",
		tonic.Handler(services.PromptRegistryService.LatestPromptForPromptId, 200))
	route.GET("/:project_id/prompt",
		tonic.Handler(services.PromptRegistryService.LatestPromptVersionsInProject, 200))

	route.POST("/:project_id/prompt",
		tonic.Handler(services.PromptRegistryService.AddNewPrompt, 200))
}

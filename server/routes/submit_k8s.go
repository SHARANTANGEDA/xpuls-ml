package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xpuls-com/xpuls-ml/deployment-core/k8s"
)

func SubmitK8s(route *gin.RouterGroup) {
	deployManager := k8s.DeployManager{}

	route.POST("/deployment", deployManager.ProcessDeployment)
}

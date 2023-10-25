package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/xpuls-com/xpuls-ml/services"
)

func AnalyticsAndUsage(route *gin.RouterGroup) {

	route.POST("/usage",
		tonic.Handler(services.AnalyticsAndUsageService.GetSqlTimeSeriesData, 200))
}

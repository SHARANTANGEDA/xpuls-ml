package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() (*gin.Engine, error) {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	SubmitK8s(router.Group("/v1/submit/k8s"))
	ProjectRoutes(router.Group("/v1/project"))
	LangChainRuns(router.Group("/"))
	return router, nil
}

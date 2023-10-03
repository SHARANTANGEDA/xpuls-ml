package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewRouter() (*gin.Engine, error) {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Configuring CORS
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:3000"}, // Allow only these origins
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))

	SubmitK8s(router.Group("/v1/submit/k8s"))
	ProjectRoutes(router.Group("/v1/project"))
	LangChainRunsTrack(router.Group("/"))
	LangChainRuns(router.Group("/v1/langchain"))
	return router, nil
}

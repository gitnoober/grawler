package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gitnoober/grawler/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/healthz", handlers.CheckHealth)
	router.POST("/crawl", handlers.CreateUrl)


	return router
}

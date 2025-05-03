package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gitnoober/grawler/handlers"
	"github.com/gitnoober/grawler/middleware"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.InjectRepositories(db))

	router.GET("/healthz", handlers.CheckHealth)
	router.POST("/url", handlers.CreateUrl)
	router.GET("/url", handlers.FetchAllUrls)
	router.POST("/crawl", handlers.CrawlUrls)
	router.GET("/task-response", handlers.FetchAllTaskResponses)

	return router
}

package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gitnoober/grawler/repository"
	"gorm.io/gorm"
)


func InjectRepositories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		urlRepo := repository.NewUrlRepository(db)
		taskRepo := repository.NewTaskRepository(db)
		taskResponseRepo := repository.NewTaskResponseRepository(db)
		urlSummaryRepo := repository.NewUrlSummaryRepository(db)
		c.Set("urlRepository", urlRepo)
		c.Set("taskRepository", taskRepo)
		c.Set("taskResponseRepository", taskResponseRepo)
		c.Set("urlSummaryRepository", urlSummaryRepo)
		fmt.Println("Injecting repositories", urlRepo, taskRepo, taskResponseRepo, urlSummaryRepo)
		c.Next()
	}
}
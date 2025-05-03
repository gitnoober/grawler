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
		
		c.Set("urlRepository", urlRepo)
		c.Set("taskRepository", taskRepo)
		fmt.Println("Injecting repositories", urlRepo, taskRepo)
		c.Next()
	}
}
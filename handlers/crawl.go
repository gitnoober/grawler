package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitnoober/grawler/models"
	"github.com/gitnoober/grawler/queue"
	"github.com/gitnoober/grawler/repository"
)

// CrawlUrls handles the creation of tasks for all existing URLs
// @Summary Create tasks for all URLs
// @Description Creates a new task for each URL in the database
// @Produce json
// @Success 201 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /crawl [post]
func CrawlUrls(c *gin.Context) {
	// Get repositories from context
	urlRepo, ok := c.MustGet("urlRepository").(repository.UrlRepository)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get URL repository"})
		return
	}

	// Get all URLs
	urls, err := urlRepo.GetAllUrls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get URLs"})
		return
	}

	// Create tasks for each URL
	for _, url := range urls {
		task := &models.Task{
			Status: models.TaskStatusPending,
			UrlID:  url.ID,
		}
		queue.CrawlQueue <- task
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "Tasks created successfully",
		"count":   len(urls),
	})
}
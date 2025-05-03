package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitnoober/grawler/repository"
)

func FetchAllTaskResponses(c *gin.Context) {
	taskResponseRepo := c.MustGet("taskResponseRepository").(repository.TaskResponseRepository)
	taskResponses, err := taskResponseRepo.GetAllTaskResponses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, taskResponses)
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitnoober/grawler/models"
	"github.com/gitnoober/grawler/repository"
)

var urlRepository repository.UrlRepository

type AddUrlRequest struct {
	Url string `json:"url" binding:"required"`
	MaxDepth int `json:"max_depth" binding:"required"`
	IsDisabled bool `json:"is_disabled"`
}

func CreateUrl(c *gin.Context){
	var request AddUrlRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := &models.Url{
		Url: request.Url,
		MaxDepth: request.MaxDepth,
		IsDisabled: request.IsDisabled,
	}
	
	if err := urlRepository.CreateUrl(url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL added successfully"})
}
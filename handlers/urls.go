package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitnoober/grawler/models"
	"github.com/gitnoober/grawler/repository"
)

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
	urlRepo, ok := c.MustGet("urlRepository").(repository.UrlRepository)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get URL repository"})
		return
	}
	if err := urlRepo.CreateUrl(url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL added successfully"})
}

func FetchAllUrls(c *gin.Context) {
	urlRepo, ok := c.MustGet("urlRepository").(repository.UrlRepository)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get URL repository"})
		return
	}
	urls, err := urlRepo.GetAllUrls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch URLs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"urls": urls})
}
package handlers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gitnoober/grawler/repository"
)

func FetchAllUrlSummaries(c *gin.Context) {
	urlSummaryRepo := c.MustGet("urlSummaryRepository").(repository.UrlSummaryRepository)
	urlSummaries, err := urlSummaryRepo.GetAllUrlSummaries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, urlSummaries)
}

func FetchUrlSummary(c *gin.Context) {
	urlSummaryRepo := c.MustGet("urlSummaryRepository").(repository.UrlSummaryRepository)
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	urlSummary, err := urlSummaryRepo.GetUrlSummary(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, urlSummary)
}

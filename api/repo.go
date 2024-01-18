package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get repositories list
// @Summary Get repositories list.
func (h *Handler) GetRepositories(c *gin.Context) {
	var uri OwnerUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hsting, exists := c.Get("hosting")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}

	repos, err := hsting.(hosting.Hosting).GetRepositories(c.Request.Context(), uri.Owner)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, repos)
}

// Get repository details.
// @Summary Get repository details.
func (h *Handler) GetRepository(c *gin.Context) {
	var uri RepoURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hsting, exists := c.Get("hosting")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}

	repo, err := hsting.(hosting.Hosting).GetRepository(c.Request.Context(), uri.Owner, uri.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, repo)
}

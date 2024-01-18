package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get repositories list
// @Summary Get repositories list.
func (h *Handler) GetRepositories(c *gin.Context) {
	var uri OwnerUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hsting, err := getHostingFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repos, err := hsting.GetRepositories(c.Request.Context(), uri.Owner)
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

	hsting, err := getHostingFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo, err := hsting.GetRepository(c.Request.Context(), uri.Owner, uri.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, repo)
}

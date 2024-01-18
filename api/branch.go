package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get branches list
// @Summary Get branches list.
func (h *Handler) GetBranches(c *gin.Context) {
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

	branches, err := hsting.(hosting.Hosting).GetBranches(c.Request.Context(), &hosting.Repository{Owner: uri.Owner, Name: uri.Name})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, branches)
}

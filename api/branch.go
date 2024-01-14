package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get branches list
// @Summary Get branches list.
func (h *Handler) GetBranches(c *gin.Context) {
	uri := RepoURI{}
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	hsting, exists := c.Get("hosting")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}

	repo := hosting.Repository{Owner: uri.Owner, Name: uri.Repo}
	branches, err := hsting.(hosting.Hosting).GetBranches(c.Request.Context(), &repo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, branches)
}

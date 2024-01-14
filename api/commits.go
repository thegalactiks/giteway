package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get commits list
// @Summary Get commits list.
func (h *Handler) GetCommits(c *gin.Context) {
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
	branches, err := hsting.(hosting.Hosting).GetCommits(c.Request.Context(), &repo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, branches)
}

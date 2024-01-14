package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get files list
// @Summary Get files list.
func (h *Handler) GetFiles(c *gin.Context) {
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
	files, err := hsting.(hosting.Hosting).GetFiles(c.Request.Context(), &repo, "/")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, files)
}

// Get file content
// @Summary Get file content.
func (h *Handler) GetFile(c *gin.Context) {
	c.JSON(200, hosting.File{})
}

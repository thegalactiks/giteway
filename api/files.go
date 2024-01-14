package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get files list
// @Summary Get files list.
func (h *Handler) GetFiles(c *gin.Context) {
	c.JSON(200, []hosting.File{})
}

// Get file content
// @Summary Get file content.
func (h *Handler) GetFile(c *gin.Context) {
	c.JSON(200, hosting.File{})
}

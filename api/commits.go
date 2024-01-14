package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get commits list
// @Summary Get commits list.
func (h *Handler) GetCommits(c *gin.Context) {
	c.JSON(200, []hosting.Commit{})
}

// Get commit by ref
// @Summary Get commit by ref.
func (h *Handler) GetCommit(c *gin.Context) {
	c.JSON(200, hosting.Commit{})
}

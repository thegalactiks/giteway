package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get branches list
// @Summary Get branches list.
func (h *Handler) GetBranches(c *gin.Context) {
	c.JSON(200, []hosting.Branch{})
}

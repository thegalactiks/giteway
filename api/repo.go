package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get repositories list
// @Summary Get repositories list.
func (h *Handler) GetRepositories(c *gin.Context) {
	c.JSON(200, []hosting.Repository{})
}

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

	hsting, err := getHostingFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branches, err := hsting.GetBranches(c.Request.Context(), &hosting.Repository{Owner: uri.Owner, Name: uri.Name})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, branches)
}

type CreateBranchForm struct {
	Name string `uri:"name" binding:"required"`
}

// Create a branch with name
// @Summary Create a branch with name.
func (h *Handler) CreateBranch(c *gin.Context) {
	var uri RepoURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var form CreateBranchForm
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hsting, err := getHostingFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branch, err := hsting.CreateBranch(
		c.Request.Context(),
		&hosting.Repository{Owner: uri.Owner, Name: uri.Name},
		&hosting.CreateBranchOpts{
			Branch: &form.Name,
		},
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, branch)
}

type DeleteBranchUri struct {
	RepoURI
	Branch string `uri:"branch" binding:"required"`
}

// Delete a branch by name
// @Summary Delete a branch by name.
func (h *Handler) DeleteBranch(c *gin.Context) {
	var uri DeleteBranchUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hsting, err := getHostingFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = hsting.DeleteBranch(
		c.Request.Context(),
		&hosting.Repository{Owner: uri.Owner, Name: uri.Name},
		&hosting.Branch{
			Name: uri.Branch,
		},
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(204)
}

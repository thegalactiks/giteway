package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get branches list
// @Summary Get branches list.
func (h *Handler) GetBranches(ctx *gin.Context) {
	var uri RepoURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	branches, err := hsting.GetBranches(ctx.Request.Context(), &hosting.Repository{Owner: uri.Owner, Name: uri.Name})
	if err != nil {
		respondWithError(ctx, http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, branches)
}

type CreateBranchForm struct {
	Name string `uri:"name" binding:"required"`
}

// Create a branch with name
// @Summary Create a branch with name.
func (h *Handler) CreateBranch(ctx *gin.Context) {
	var uri RepoURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	var form CreateBranchForm
	if err := ctx.ShouldBind(&form); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	branch, err := hsting.CreateBranch(
		ctx.Request.Context(),
		&hosting.Repository{Owner: uri.Owner, Name: uri.Name},
		&hosting.CreateBranchOpts{
			Branch: &form.Name,
		},
	)
	if err != nil {
		respondWithError(ctx, http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, branch)
}

type DeleteBranchUri struct {
	RepoURI
	Branch string `uri:"branch" binding:"required"`
}

// Delete a branch by name
// @Summary Delete a branch by name.
func (h *Handler) DeleteBranch(ctx *gin.Context) {
	var uri DeleteBranchUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	err = hsting.DeleteBranch(
		ctx.Request.Context(),
		&hosting.Repository{Owner: uri.Owner, Name: uri.Name},
		&hosting.Branch{
			Name: uri.Branch,
		},
	)
	if err != nil {
		respondWithError(ctx, http.StatusBadGateway, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

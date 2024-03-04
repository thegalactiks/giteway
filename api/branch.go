package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

func (h *Handler) GetBranches(ctx *gin.Context) {
	hostingService := ctx.MustGet("hosting").(hosting.GitHostingService)
	repo := ctx.MustGet("repo").(*hosting.Repository)

	branches, err := hostingService.GetBranches(ctx.Request.Context(), repo)
	if err != nil {
		RespondError(ctx, http.StatusBadRequest, "failed to get branches", err)
		return
	}

	RespondJSON(ctx, http.StatusOK, branches)
}

type CreateBranchForm struct {
	Name string `uri:"name" binding:"required"`
}

func (h *Handler) CreateBranch(ctx *gin.Context) {
	hostingService := ctx.MustGet("hosting").(hosting.GitHostingService)
	repo := ctx.MustGet("repo").(*hosting.Repository)

	var form CreateBranchForm
	if err := ctx.ShouldBind(&form); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	branch, err := hostingService.CreateBranch(
		ctx.Request.Context(),
		repo,
		&hosting.CreateBranchOpts{
			Branch: &form.Name,
		},
	)
	if err != nil {
		RespondError(ctx, http.StatusBadRequest, "failed to create branch", err)
		return
	}

	RespondJSON(ctx, http.StatusCreated, branch)
}

type DeleteBranchUri struct {
	Branch string `uri:"branch" binding:"required"`
}

func (h *Handler) DeleteBranch(ctx *gin.Context) {
	var uri DeleteBranchUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	hostingService := ctx.MustGet("hosting").(hosting.GitHostingService)
	repo := ctx.MustGet("repo").(*hosting.Repository)

	err := hostingService.DeleteBranch(
		ctx.Request.Context(),
		repo,
		&hosting.Branch{
			Name: uri.Branch,
		},
	)
	if err != nil {
		RespondError(ctx, http.StatusBadRequest, "failed to delete branch", err)
		return
	}

	RespondJSON(ctx, http.StatusNoContent, nil)
}

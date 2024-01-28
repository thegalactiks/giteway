package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get repositories list
// @Summary Get repositories list.
func (h *Handler) GetRepositories(ctx *gin.Context) {
	var uri OwnerUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		WriteErr(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		WriteErr(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	repos, err := hsting.GetRepositories(ctx.Request.Context(), uri.Owner)
	if err != nil {
		WriteErr(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
	}

	ctx.JSON(http.StatusOK, repos)
}

// Get repository details.
// @Summary Get repository details.
func (h *Handler) GetRepository(ctx *gin.Context) {
	var uri RepoURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		WriteErr(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		WriteErr(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	repo, err := hsting.GetRepository(ctx.Request.Context(), uri.Owner, uri.Name)
	if err != nil {
		WriteErr(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	ctx.JSON(http.StatusOK, repo)
}

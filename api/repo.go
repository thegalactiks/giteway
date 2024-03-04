package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

func (h *Handler) GetRepositories(ctx *gin.Context) {
	hostingService := ctx.MustGet("hosting").(hosting.GitHostingService)
	owner := ctx.MustGet("owner").(string)

	repos, err := hostingService.GetRepositories(ctx.Request.Context(), owner)
	if err != nil {
		RespondError(ctx, http.StatusBadRequest, "failed to get repositories", err)
		return
	}

	RespondJSON(ctx, http.StatusOK, repos)
}

func (h *Handler) GetRepository(ctx *gin.Context) {
	hostingService := ctx.MustGet("hosting").(hosting.GitHostingService)
	repo := ctx.MustGet("repo").(*hosting.Repository)

	repo, err := hostingService.GetRepository(ctx.Request.Context(), repo.Owner, repo.Name)
	if err != nil {
		RespondError(ctx, http.StatusBadRequest, "failed to get repository", err)
		return
	}

	RespondJSON(ctx, http.StatusOK, repo)
}

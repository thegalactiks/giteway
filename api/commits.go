package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

func (h *Handler) GetCommits(ctx *gin.Context) {
	hostingService := ctx.MustGet("hosting").(hosting.GitHostingService)
	repo := ctx.MustGet("repo").(*hosting.Repository)

	var form RefForm
	if err := ctx.ShouldBind(&form); err != nil {
		RespondError(ctx, http.StatusBadRequest, "failed to bind form", err)
		return
	}

	commits, err := hostingService.GetCommits(
		ctx.Request.Context(),
		repo,
		&hosting.GetCommitsOpts{Ref: form.Ref},
	)
	if err != nil {
		RespondError(ctx, http.StatusBadRequest, "failed to get commits", err)
		return
	}

	RespondJSON(ctx, http.StatusOK, commits)
}

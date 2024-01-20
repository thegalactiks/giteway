package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

// Get commits list
// @Summary Get commits list.
func (h *Handler) GetCommits(ctx *gin.Context) {
	var uri RepoURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	var form RefForm
	if err := ctx.ShouldBind(&form); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	commits, err := hsting.GetCommits(
		ctx.Request.Context(),
		&hosting.Repository{Owner: uri.Owner, Name: uri.Name},
		&hosting.GetCommitsOpts{Ref: form.Ref},
	)
	if err != nil {
		respondWithError(ctx, http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, commits)
}

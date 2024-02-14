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
		WriteErr(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	var form RefForm
	if err := ctx.ShouldBind(&form); err != nil {
		WriteErr(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		WriteErr(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	commits, err := hsting.GetCommits(
		ctx.Request.Context(),
		&hosting.Repository{Owner: uri.Owner, Name: uri.Name},
		&hosting.GetCommitsOpts{Ref: form.Ref},
	)
	if err != nil {
		WriteErr(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	ctx.JSON(http.StatusOK, commits)
}

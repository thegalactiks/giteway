package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

type FileURI struct {
	Path string `uri:"path,default=/"`
}

type CommitForm struct {
	Message *string `form:"message,omitempty" json:"message,omitempty"`
}

type FileContentForm struct {
	Encoding string     `form:"encoding,default=text" json:"encoding"`
	Content  string     `form:"content" json:"content" binding:"required"`
	Commit   CommitForm `form:"commit,omitempty" json:"commit,omitempty"`
}

func (h *Handler) GetFiles(ctx *gin.Context) {
	hostingService := ctx.MustGet("hosting").(hosting.GitHostingService)
	repo := ctx.MustGet("repo").(*hosting.Repository)

	var uri FileURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	var form RefForm
	if err := ctx.ShouldBind(&form); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	switch ctx.NegotiateFormat(gin.MIMEJSON, RawMimeTypes) {
	case RawMimeTypes:
		file, err := hostingService.GetRawFile(ctx.Request.Context(), repo, uri.Path, &hosting.GetFileOpts{})
		if err != nil {
			RespondError(ctx, http.StatusBadRequest, "failed to get raw file", err)
			return
		}

		ctx.Data(http.StatusOK, RawMimeTypes, file)
		return

	default:
		file, files, err := hostingService.GetFiles(ctx.Request.Context(), repo, uri.Path)
		if err != nil {
			RespondError(ctx, http.StatusBadRequest, "failed to get files", err)
			return
		}

		if file != nil {
			RespondJSON(ctx, http.StatusOK, file)
			return
		}

		RespondJSON(ctx, http.StatusOK, files)
	}
}

func (h *Handler) CreateFile(ctx *gin.Context) {
	hostingService := ctx.MustGet("hosting").(hosting.GitHostingService)
	repo := ctx.MustGet("repo").(*hosting.Repository)

	var uri FileURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	var queryForm RefForm
	if err := ctx.ShouldBindQuery(&queryForm); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	var form FileContentForm
	if err := ctx.ShouldBind(&form); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	var message string
	if form.Commit.Message != nil {
		message = *form.Commit.Message
	} else {
		message = "chore: create file"
	}

	file := hosting.File{
		Path:    uri.Path,
		Content: &form.Content,
	}
	file.SetEncoding(form.Encoding)
	_, commit, err := hostingService.CreateFile(ctx, repo, &file, &hosting.CreateFileOpts{
		SHA:    queryForm.SHA,
		Branch: &queryForm.Branch,
		Ref:    queryForm.Ref,
		Commit: hosting.Commit{
			Message: &message,
		},
	})
	if err != nil {
		RespondError(ctx, http.StatusBadRequest, "failed to create file", err)
		return
	}

	ctx.JSON(http.StatusCreated, commit)
}

func (h *Handler) UpdateFile(ctx *gin.Context) {
	hostingService := ctx.MustGet("hosting").(hosting.GitHostingService)
	repo := ctx.MustGet("repo").(*hosting.Repository)

	var uri FileURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	var queryForm RefForm
	if err := ctx.ShouldBindQuery(&queryForm); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	var form FileContentForm
	if err := ctx.ShouldBind(&form); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	var message string
	if form.Commit.Message != nil {
		message = *form.Commit.Message
	} else {
		message = "chore: update file"
	}

	file := hosting.File{
		Path:    uri.Path,
		Content: &form.Content,
	}
	file.SetEncoding(form.Encoding)
	_, commit, err := hostingService.UpdateFile(ctx, repo, &file, &hosting.UpdateFileOpts{
		SHA:    queryForm.SHA,
		Branch: &queryForm.Branch,
		Ref:    queryForm.Ref,
		Commit: hosting.Commit{
			Message: &message,
		},
	})
	if err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	ctx.JSON(http.StatusOK, commit)
}

func (h *Handler) DeleteFile(ctx *gin.Context) {
	hostingService := ctx.MustGet("hosting").(hosting.GitHostingService)
	repo := ctx.MustGet("repo").(*hosting.Repository)

	var uri FileURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	var queryForm RefForm
	if err := ctx.ShouldBindQuery(&queryForm); err != nil {
		RespondError(ctx, http.StatusBadRequest, HTTPRequestValidationFailed, err)
		return
	}

	message := "chore: delete file"
	_, err := hostingService.DeleteFile(ctx, repo, uri.Path, &hosting.DeleteFileOpts{
		SHA:    queryForm.SHA,
		Branch: &queryForm.Branch,
		Ref:    queryForm.Ref,
		Commit: hosting.Commit{
			Message: &message,
		},
	})
	if err != nil {
		RespondError(ctx, http.StatusBadRequest, "failed to delete file", err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

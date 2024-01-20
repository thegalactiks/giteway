package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

type FileURI struct {
	RepoURI
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

// Get files list
// @Summary Get files list.
func (h *Handler) GetFiles(ctx *gin.Context) {
	var uri FileURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	var form RefForm
	if err := ctx.ShouldBind(&form); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	repo := hosting.Repository{Owner: uri.Owner, Name: uri.Name}
	switch ctx.NegotiateFormat(gin.MIMEJSON, RawMimeTypes) {
	case RawMimeTypes:
		file, err := hsting.GetRawFile(ctx.Request.Context(), &repo, uri.Path, &hosting.GetFileOpts{})
		if err != nil {
			respondWithError(ctx, http.StatusBadGateway, err)
			return
		}

		ctx.Data(http.StatusOK, RawMimeTypes, file)
		return

	default:
		file, files, err := hsting.GetFiles(ctx.Request.Context(), &repo, uri.Path)
		if err != nil {
			respondWithError(ctx, http.StatusBadGateway, err)
			return
		}

		if file != nil {
			ctx.JSON(http.StatusOK, file)
			return
		}

		ctx.JSON(http.StatusOK, files)
	}
}

// Create file
// @Summary Create file.
func (h *Handler) CreateFile(ctx *gin.Context) {
	var uri FileURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	var form FileContentForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	var queryForm RefForm
	if err := ctx.ShouldBindQuery(&queryForm); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	var message string
	if form.Commit.Message != nil {
		message = *form.Commit.Message
	} else {
		message = "chore: create file"
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	file := hosting.File{
		Path:    uri.Path,
		Content: &form.Content,
	}
	file.SetEncoding(form.Encoding)
	_, commit, err := hsting.CreateFile(ctx, &hosting.Repository{Owner: uri.Owner, Name: uri.Name}, &file, &hosting.CreateFileOpts{
		SHA:    queryForm.SHA,
		Branch: &queryForm.Branch,
		Ref:    queryForm.Ref,
		Commit: hosting.Commit{
			Message: message,
		},
	})
	if err != nil {
		respondWithError(ctx, http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, commit)
}

// Update file
// @Summary Update file.
func (h *Handler) UpdateFile(ctx *gin.Context) {
	var uri FileURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	var form FileContentForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	var queryForm RefForm
	if err := ctx.ShouldBindQuery(&queryForm); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	var message string
	if form.Commit.Message != nil {
		message = *form.Commit.Message
	} else {
		message = "chore: update file"
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	file := hosting.File{
		Path:    uri.Path,
		Content: &form.Content,
	}
	file.SetEncoding(form.Encoding)
	_, commit, err := hsting.UpdateFile(ctx, &hosting.Repository{Owner: uri.Owner, Name: uri.Name}, &file, &hosting.UpdateFileOpts{
		SHA:    queryForm.SHA,
		Branch: &queryForm.Branch,
		Ref:    queryForm.Ref,
		Commit: hosting.Commit{
			Message: message,
		},
	})
	if err != nil {
		respondWithError(ctx, http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, commit)
}

// Delete file
// @Summary Delete file.
func (h *Handler) DeleteFile(ctx *gin.Context) {
	var uri FileURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	var queryForm RefForm
	if err := ctx.ShouldBindQuery(&queryForm); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	hsting, err := getHostingFromContext(ctx)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	message := "chore: delete file"
	_, err = hsting.DeleteFile(ctx, &hosting.Repository{Owner: uri.Owner, Name: uri.Name}, uri.Path, &hosting.DeleteFileOpts{
		SHA:    queryForm.SHA,
		Branch: &queryForm.Branch,
		Ref:    queryForm.Ref,
		Commit: hosting.Commit{
			Message: message,
		},
	})
	if err != nil {
		respondWithError(ctx, http.StatusBadGateway, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

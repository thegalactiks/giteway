package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
)

type FileURI struct {
	RepoURI
	Path string `json:"path" uri:"path,default=/"`
}

// Get files list
// @Summary Get files list.
func (h *Handler) GetFiles(c *gin.Context) {
	uri := FileURI{}
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	hsting, exists := c.Get("hosting")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}

	repo := hosting.Repository{Owner: uri.Owner, Name: uri.Repo}
	switch c.NegotiateFormat(gin.MIMEJSON, RawMimeTypes) {
	case RawMimeTypes:
		file, err := hsting.(hosting.Hosting).GetRawFile(c.Request.Context(), &repo, uri.Path)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Data(200, RawMimeTypes, file)
		return

	default:
		file, files, err := hsting.(hosting.Hosting).GetFiles(c.Request.Context(), &repo, uri.Path)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		if file != nil {
			c.JSON(200, file)
			return
		}

		c.JSON(200, files)
	}
}

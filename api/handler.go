package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/api/middlewares"
	"github.com/thegalactiks/giteway/internal/config"
)

const (
	RawMimeTypes string = "application/vnd.giteway.raw"
)

type Handler struct {
	c *config.Config
	e *gin.Engine
}

type OwnerUri struct {
	Hosting string `uri:"hosting" binding:"required"`
	Owner   string `uri:"owner" binding:"required"`
}

type RepoURI struct {
	OwnerUri
	Name string `uri:"name" binding:"required"`
}

type RefForm struct {
	Ref *string `form:"ref,default=main"`
}

func NewHandler(c *config.Config, e *gin.Engine) *Handler {
	return &Handler{
		c: c,
		e: e,
	}
}

func Routes(r *gin.Engine, h *Handler) {
	adminApi := r.Group("/")
	adminApi.Use(middlewares.HostingMiddleware)
	adminApi.GET("/repos/:hosting/:owner", h.GetRepositories)
	adminApi.GET("/repos/:hosting/:owner/:name", h.GetRepository)

	gitRepoApi := r.Group("/repos/:hosting/:owner/:name")

	gitRepoApi.Use(middlewares.HostingMiddleware)
	gitRepoApi.GET("/branches", h.GetBranches)
	gitRepoApi.GET("/commits", h.GetCommits)
	gitRepoApi.GET("/files", h.GetFiles)
	gitRepoApi.GET("/files/*path", h.GetFiles)
}

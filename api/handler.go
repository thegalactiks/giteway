package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/internal/config"
)

const (
	RawMimeTypes string = "application/vnd.giteway.raw"
)

type Handler struct {
	c *config.Config
	e *gin.Engine
}

type RefForm struct {
	Ref    *string `form:"ref,omitempty"`
	SHA    *string `form:"sha,omitempty"`
	Branch string  `form:"branch,default=main"`
}

func NewHandler(cfg *config.Config, e *gin.Engine) *Handler {
	return &Handler{cfg, e}
}

func Routes(r *gin.Engine, h *Handler) {
	adminApi := r.Group("/")
	adminApi.Use(HostingMiddleware(), OwnerMiddleware())
	adminApi.GET("/repos/:hosting/:owner", h.GetRepositories)
	adminApi.GET("/repos/:hosting/:owner/:repo", RepoMiddleware(), h.GetRepository)

	gitRepoApi := adminApi.Group("/repos/:hosting/:owner/:repo")
	gitRepoApi.Use(RepoMiddleware())
	gitRepoApi.GET("/branches", h.GetBranches)
	gitRepoApi.POST("/branches", h.CreateBranch)
	gitRepoApi.DELETE("/branches/:branch", h.DeleteBranch)
	gitRepoApi.GET("/commits", h.GetCommits)
	gitRepoApi.GET("/files", h.GetFiles)
	gitRepoApi.GET("/files/*path", h.GetFiles)
	gitRepoApi.POST("/files/*path", h.CreateFile)
	gitRepoApi.PUT("/files/*path", h.UpdateFile)
	gitRepoApi.DELETE("/files/*path", h.DeleteFile)
}

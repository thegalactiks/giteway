package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/api/middlewares"
	"github.com/thegalactiks/giteway/internal/config"
)

type Handler struct {
	c *config.Config
	e *gin.Engine
}

type RepoURI struct {
	Hosting string `json:"hosting" uri:"hosting"`
	Owner   string `json:"owner" uri:"owner"`
	Repo    string `json:"repo" uri:"repo"`
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
	adminApi.GET("/repos/:hosting/:owner/:repo", h.GetRepository)

	gitRepoApi := r.Group("/repos/:hosting/:owner/:repo")

	gitRepoApi.Use(middlewares.HostingMiddleware)
	gitRepoApi.GET("/branches", h.GetBranches)
	// gitRepoApi.GET("/commits", h.GetCommits)
	// gitRepoApi.GET("/commits/:ref", h.GetCommit)
	// gitRepoApi.GET("/files", h.GetFiles)
	// gitRepoApi.GET("/files/:path", h.GetFile)
}

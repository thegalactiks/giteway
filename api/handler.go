package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/internal/config"
)

type Handler struct {
	c *config.Config
	r *gin.Engine
}

func NewHandler(c *config.Config, r *gin.Engine) *Handler {
	return &Handler{
		c: c,
		r: r,
	}
}

func Routes(r *gin.Engine, h *Handler) {
	adminApi := r.Group("/")

	adminApi.GET("/repos", h.GetRepositories)

	gitRepoApi := r.Group("/repos/:hosting/:owner/:repo")

	gitRepoApi.GET("/branches", h.GetBranches)
	gitRepoApi.GET("/commits", h.GetCommits)
	gitRepoApi.GET("/commits/:ref", h.GetCommit)
	gitRepoApi.GET("/files", h.GetFiles)
	gitRepoApi.GET("/files/:path", h.GetFile)
}

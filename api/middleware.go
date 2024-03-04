package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/github"
	"github.com/thegalactiks/giteway/gitlab"
	"github.com/thegalactiks/giteway/hosting"
)

func getTokenFromContext(ctx *gin.Context) (*string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return nil, nil
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return nil, errors.New("invalid authorization header")
	}

	token := authParts[1]

	return &token, nil
}

func getHostingFromContext(hosting string, token *string) (hosting.GitHostingService, error) {
	switch hosting {
	case "github.com":
		return github.NewGithubService(token)

	case "gitlab.com":
		if token == nil || *token == "" {
			return nil, errors.New("gitlab require a token")
		}

		return gitlab.NewGitlabService(*token)
	}

	return nil, errors.New("unknown hosting service")
}

func HostingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getTokenFromContext(ctx)
		if err != nil {
			RespondError(ctx, http.StatusUnauthorized, "invalid authorization header", err)
			ctx.Abort()
			return
		}

		hosting := ctx.Param("hosting")
		service, err := getHostingFromContext(hosting, token)
		if err != nil {
			RespondError(ctx, http.StatusBadRequest, "unknown git provider", err)
			ctx.Abort()
			return
		}

		ctx.Set("hosting", service)
		ctx.Next()
	}
}

func OwnerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		owner := ctx.Param("owner")

		if owner == "" {
			RespondError(ctx, http.StatusBadRequest, "invalid owner uri")
			ctx.Abort()
			return
		}

		ctx.Set("owner", owner)
		ctx.Next()
	}
}

func RepoMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		repoParam := ctx.Param("repo")
		if repoParam == "" {
			RespondError(ctx, http.StatusBadRequest, "invalid repository uri")
			ctx.Abort()
			return
		}

		repo := &hosting.Repository{
			Owner: ctx.MustGet("owner").(string),
			Name:  repoParam,
		}

		ctx.Set("repo", repo)
		ctx.Next()
	}
}

func PathMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Param("path")

		if path == "" {
			RespondError(ctx, http.StatusBadRequest, "invalid path uri")
			ctx.Abort()
			return
		}

		ctx.Set("path", path)
		ctx.Next()
	}
}

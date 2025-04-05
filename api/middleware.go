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

const (
	GithubHost = "github.com"
	GitlabHost = "gitlab.com"
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
	if token == "" {
		return nil, errors.New("invalid authorization header")
	}

	return &token, nil
}

func HostingMiddleware(githubService *github.GithubService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getTokenFromContext(ctx)
		if err != nil {
			RespondError(ctx, http.StatusUnauthorized, "invalid authorization header", err)
			ctx.Abort()
			return
		}

		hostingParam := ctx.Param("hosting")
		ownerParam := ctx.Param("owner")
		var service hosting.GitHostingService
		switch hostingParam {
		case GithubHost:
			switch {
			case token != nil:
				service, err = githubService.WithAuthToken(ctx, *token)
				if err != nil {
					RespondError(ctx, http.StatusUnauthorized, "invalid github token", err)
					ctx.Abort()
					return
				}
			case ownerParam != "" && githubService.IsKnownInstallation(ownerParam):
				service, err = githubService.WithInstallationOwner(ownerParam)
				if err != nil {
					RespondError(ctx, http.StatusUnauthorized, "invalid github installation", err)
					ctx.Abort()
					return
				}
			default:
				service = githubService
			}
		case GitlabHost:
			service, err = gitlab.NewGitlabService(*token)
			if err != nil {
				RespondError(ctx, http.StatusUnauthorized, "invalid gitlab token", err)
				ctx.Abort()
				return
			}
		default:
			RespondError(ctx, http.StatusBadRequest, "unknown git provider")
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

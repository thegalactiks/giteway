package api

import (
	"errors"
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

func getHostingFromContext(ctx *gin.Context) (hosting.GitHostingService, error) {
	token, err := getTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	switch ctx.Param("hosting") {
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

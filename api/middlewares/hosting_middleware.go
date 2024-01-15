package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
	"github.com/thegalactiks/giteway/internal/hosting/github"
	"github.com/thegalactiks/giteway/internal/hosting/gitlab"
)

func getTokenFromContext(c *gin.Context) (*string, error) {
	authHeader := c.GetHeader("Authorization")
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

func getHostingFromURLParam(c *gin.Context) (hosting.Hosting, error) {
	switch c.Param("hosting") {
	case "github.com":
		token, err := getTokenFromContext(c)
		if err != nil {
			return nil, err
		}

		return github.New(token)

	case "gitlab.com":
		token, err := getTokenFromContext(c)
		if err != nil {
			return nil, err
		}

		if token == nil || *token == "" {
			return nil, errors.New("gitlab require a token")
		}

		return gitlab.New(*token)
	}

	return nil, errors.New("unknown hosting service")
}

func HostingMiddleware(c *gin.Context) {
	h, err := getHostingFromURLParam(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Set("hosting", h)
	c.Next()
}

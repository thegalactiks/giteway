package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/thegalactiks/giteway/hosting"
	"github.com/thegalactiks/giteway/internal/hosting/github"
	"github.com/thegalactiks/giteway/internal/hosting/gitlab"
)

func handleValidationErrors(c *gin.Context, errs validator.ValidationErrors) {
	errorMessages := make([]string, len(errs))
	for i, err := range errs {
		errorMessages[i] = fmt.Sprintf("%s is %s", err.Field(), err.Tag())
	}

	c.JSON(http.StatusBadRequest, gin.H{"validation_error": errorMessages})
}

func respondWithError(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{"error": err.Error()})
}

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

func getHostingFromContext(ctx *gin.Context) (hosting.Hosting, error) {
	token, err := getTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	switch ctx.Param("hosting") {
	case "github.com":
		return github.New(token)

	case "gitlab.com":
		if token == nil || *token == "" {
			return nil, errors.New("gitlab require a token")
		}

		return gitlab.New(*token)
	}

	return nil, errors.New("unknown hosting service")
}

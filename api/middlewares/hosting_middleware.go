package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegalactiks/giteway/hosting"
	"github.com/thegalactiks/giteway/internal/hosting/github"
)

func getHostingFromURLParam(hostingParam string) hosting.Hosting {
	switch hostingParam {
	case "github.com":
		return github.New()
	}

	return nil
}

func HostingMiddleware(c *gin.Context) {
	h := getHostingFromURLParam(c.Param("hosting"))
	if h == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "unknown hosting service")
		return
	}

	c.Set("hosting", h)
	c.Next()
}

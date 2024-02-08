package healthz

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"project.com/restful-api/repositories"
)

func GetStatus(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelWarn, "Get Healthz.GetStatus")
	c.Header("cache-control", "no-cache")
	c.Status(http.StatusOK)

	bodysize := c.Request.ContentLength
	paramsize := len(c.Request.URL.Query())

	if bodysize > 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if paramsize > 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var repo = repositories.GetHealthzRepo()
	err := repo.GetStatus()
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Healthz endpoint is down")
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}
}

package healthz

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"project.com/restful-api/services/healthz"
	"project.com/restful-api/utilities"
)

func InitController(router *gin.Engine) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Setting up health controller")
	router.Any("/healthz", handleHealthz)
}

func handleHealthz(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Healthz endpoint")
	statserr := utilities.StatsdClient.Inc("endpoint.healthz.status", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	c.Header("cache-control", "no-cache")

	if c.Request.Method != http.MethodGet {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Healthz.Get")
		c.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}

	healthz.GetStatus(c)
}

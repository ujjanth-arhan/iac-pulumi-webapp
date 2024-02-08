package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"project.com/restful-api/controllers/assignments"
	"project.com/restful-api/controllers/healthz"
	"project.com/restful-api/services"
	"project.com/restful-api/utilities"
)

func InitControllers(router *gin.Engine) {
	// Initialize services
	services.Init()

	// Initialize all the controllers
	healthz.InitController(router)
	assignments.InitController(router)

	router.NoRoute(noroute)
}

func noroute(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Setting up handle for no route")
	statserr := utilities.StatsdClient.Inc("endpoint.generic.noroute", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	c.AbortWithStatus(http.StatusNotFound)
	return
}

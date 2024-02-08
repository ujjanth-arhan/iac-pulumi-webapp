package assignments

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"project.com/restful-api/services/assignments"
)

func InitController(router *gin.Engine) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Setting up assignments controller")
	router.Handle(http.MethodGet, "/v3/assignments/:id", assignments.Get)
	router.Handle(http.MethodGet, "/v3/assignments", assignments.GetAll)
	router.Handle(http.MethodPost, "/v3/assignments", assignments.Post)
	router.Handle(http.MethodPut, "/v3/assignments/:id", assignments.Put)
	router.Handle(http.MethodDelete, "/v3/assignments/:id", assignments.Del)
	router.Handle(http.MethodPatch, "/v3/assignments/:id", assignments.Patch)
	router.Handle(http.MethodPost, "/v3/assignments/:id/submission", assignments.PostSubmission)
	// router.Handle(http.MethodGet, "/v1/assignments/:id/submission", assignments.GetSubmissions)
}

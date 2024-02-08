package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"project.com/restful-api/constants"
	"project.com/restful-api/utilities"

	"github.com/gin-gonic/gin"

	"project.com/restful-api/controllers"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Explicit environment file not provided, will attempt to use the injected version")
	}

	utilities.LoadLogger()
	utilities.InitializeStatsd()
}

func main() {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Running in "+os.Getenv("ENVIRONMENT")+" environment")
	//gin.SetMode("release")

	router := gin.Default()

	// Initialize controllers
	controllers.InitControllers(router)
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Resources have been initialized")

	err := router.Run(":8080")
	if err != nil {
		slog.LogAttrs(context.Background(), constants.ErrorLevelFatal, "Unable to start application on localhost:8080")
	}
}

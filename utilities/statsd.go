package utilities

import (
	"context"
	"github.com/cactus/go-statsd-client/v5/statsd"
	"log/slog"
	"os"
	"project.com/restful-api/constants"
)

var StatsdClient statsd.Statter

func InitializeStatsd() {
	slog.LogAttrs(context.Background(), constants.ErrorLevelTrace, "Initializing Statsd")
	config := &statsd.ClientConfig{
		Address: os.Getenv("STATSD_SERVER"),
		Prefix:  os.Getenv("STATSD_PREFIX"),
	}

	client, err := statsd.NewClientWithConfig(config)
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Failed to configure statsd")
	}

	StatsdClient = client

	slog.LogAttrs(context.Background(), slog.LevelWarn, "Statsd has been enabled")

	// Todo: Where should statsd be closed to avoid memory leak
	//defer client.Close()
}

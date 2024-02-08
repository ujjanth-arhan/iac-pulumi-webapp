package utilities

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"project.com/restful-api/constants"
)

func LoadLogger() {
	var LevelNames = map[slog.Leveler]string{
		constants.ErrorLevelTrace: "TRACE",
		constants.ErrorLevelFatal: "FATAL",
	}

	logLevel := &slog.LevelVar{}
	logLevel.Set(slog.LevelInfo)
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}

				a.Value = slog.StringValue(levelLabel)
			}

			return a
		},
	}

	file, err := os.Create(os.Getenv("LOG_FILE"))
	if err != nil {
		fmt.Errorf("Could not create log file")
	}

	logger := slog.New(slog.NewJSONHandler(file, &opts))
	slog.SetDefault(logger)
	slog.LogAttrs(context.Background(), slog.LevelWarn, "Logging has been enabled")
}

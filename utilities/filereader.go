package utilities

import (
	"context"
	"encoding/csv"
	"log/slog"
	"os"
)

func ReadCSV(file string) [][]string {
	StatsdClient.Inc("utilities.filereader", 1, 1.0)
	f, err := os.Open(file)
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Unable open file: "+file)
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Failed to read file: "+file)
	}

	return lines
}

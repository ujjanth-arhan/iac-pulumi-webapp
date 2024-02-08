package healthz

import (
	"context"
	"gorm.io/gorm"
	"log/slog"
	"project.com/restful-api/utilities"
)

type Testconnection struct {
	Db *gorm.DB
}

func (con Testconnection) GetStatus() error {
	statserr := utilities.StatsdClient.Inc("repo.healthz.mock.GetStatus", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	return nil
}

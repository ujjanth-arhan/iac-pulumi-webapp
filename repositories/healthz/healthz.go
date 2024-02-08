package healthz

import (
	"context"
	"log/slog"
	"os"
	"project.com/restful-api/utilities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository interface {
	GetStatus() error
}

type Prodconnection struct {
	Db *gorm.DB
}

func (con Prodconnection) GetStatus() error {
	statserr := utilities.StatsdClient.Inc("repo.healthz.GetStatus", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	db, err := gorm.Open(postgres.Open(string(os.Getenv("DB_CONNECTION"))), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	pgrdb, _ := db.DB()
	pgrdb.Close()

	return err
}

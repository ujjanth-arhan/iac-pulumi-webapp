package intermediates

import (
	"context"
	"gorm.io/gorm"
	"log/slog"
	"project.com/restful-api/repositories/models"
	"project.com/restful-api/utilities"
)

type Repository interface {
	GetByAssignment(id string) models.Account_Assignment
}

type Prodconnection struct {
	Db *gorm.DB
}

func (con Prodconnection) GetByAssignment(id string) models.Account_Assignment {
	statserr := utilities.StatsdClient.Inc("repo.account_assignments.GetByAssignment", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	var aa models.Account_Assignment
	con.Db.Where("assignment_id = ?", id).Limit(1).Find(&aa)
	return aa
}

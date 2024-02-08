package submissions

import (
	"context"
	"gorm.io/gorm"
	"log/slog"
	"project.com/restful-api/repositories/models"
	"project.com/restful-api/utilities"
)

type Repository interface {
	GetByFilter(accountId string, assignmentId string) []models.Submission
	Add(submission models.Submission)
	GetByAssignmentId(accountId string) []models.Submission
}

type Prodconnection struct {
	Db *gorm.DB
}

func (con Prodconnection) GetByFilter(accountId string, assignmentId string) []models.Submission {
	statserr := utilities.StatsdClient.Inc("repo.submissions.GetByFilter", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	var submissions []models.Submission
	con.Db.Where(" account_id = ? AND assignment_id = ?", accountId, assignmentId).Find(&submissions)
	return submissions
}

func (con Prodconnection) Add(submission models.Submission) {
	statserr := utilities.StatsdClient.Inc("repo.submission.Add", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	con.Db.Create(&submission)
}

func (con Prodconnection) GetByAssignmentId(accountId string) []models.Submission {
	statserr := utilities.StatsdClient.Inc("repo.submission.Get", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	var submissions []models.Submission
	con.Db.Where(&submissions, "assignment_id = ?", accountId)

	return submissions
}

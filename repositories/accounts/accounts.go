package accounts

import (
	"context"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"project.com/restful-api/repositories/models"
	"project.com/restful-api/utilities"
)

type Repository interface {
	GetByEmail(email string) ([]models.Account, error)
	Get(id string) ([]models.Account, error)
}

type Prodconnection struct {
	Db *gorm.DB
}

func (con Prodconnection) GetByEmail(email string) ([]models.Account, error) {
	statserr := utilities.StatsdClient.Inc("repo.accounts.GetByEmail", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	var accounts []models.Account
	res := con.Db.Where("email = ?", email).Limit(1).Find(&accounts)
	if res.Error != nil {
		log.Println("Error fetching account: " + res.Error.Error())
	}

	return accounts, res.Error
}

func (con Prodconnection) Add(account models.Account) {
	statserr := utilities.StatsdClient.Inc("repo.accounts.Add", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	res := con.Db.Create(&account)
	if res.Error != nil {
		log.Println("Unable to add account")
	}
}

func (con Prodconnection) Get(id string) ([]models.Account, error) {
	statserr := utilities.StatsdClient.Inc("repo.accounts.Get", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	var accounts []models.Account
	res := con.Db.Where("id = ?", id).Limit(1).Find(&accounts)
	if res.Error != nil {
		log.Println("Error fetching account: " + res.Error.Error())
	}

	return accounts, res.Error
}

package accounts

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"project.com/restful-api/repositories"
	"project.com/restful-api/utilities"
)

func IsBasicAuthorized(header string) (string, bool) {
	utilities.StatsdClient.Inc("accounts.IsBasicAuthorized", 1, 1.0)
	creds, err := utilities.DecodeBasicAuth(header)
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Failed to decode basic auth: "+err.Error())
		return "", false
	}

	repo := repositories.GetAccountsRepo()
	accounts, err := repo.GetByEmail(creds[0])
	if err != nil || len(accounts) == 0 {
		return "", false
	}

	er := bcrypt.CompareHashAndPassword([]byte(accounts[0].Password), []byte(creds[1]))
	if er != nil {
		return "", false
	}

	return accounts[0].ID, true
}

func Get(id string) string {
	utilities.StatsdClient.Inc("accounts.Get", 1, 1.0)
	repo := repositories.GetAccountsRepo()
	accounts, err := repo.Get(id)
	if err != nil || len(accounts) == 0 {
		slog.LogAttrs(context.Background(), slog.LevelError, "Account fetch error or no account found")
	}

	return accounts[0].Email
}

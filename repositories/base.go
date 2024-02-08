package repositories

import (
	"context"
	"log/slog"
	"os"
	"project.com/restful-api/constants"
	"project.com/restful-api/repositories/submissions"
	"strconv"

	uuid2 "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"project.com/restful-api/repositories/accounts"
	"project.com/restful-api/repositories/assignments"
	"project.com/restful-api/repositories/healthz"
	"project.com/restful-api/repositories/intermediates"
	"project.com/restful-api/repositories/models"
	"project.com/restful-api/utilities"
)

var db *gorm.DB

func GetHealthzRepo() healthz.Repository {
	var env = os.Getenv("ENVIRONMENT")
	if env == "PRODUCTION" {
		return healthz.Prodconnection{Db: db}
	} else if env == "TESTING" {
		return healthz.Testconnection{Db: db}
	}

	/** OnHold Todo: Use else instead of TESTING */

	return nil
}

func GetAccountsRepo() accounts.Repository {
	var env = os.Getenv("ENVIRONMENT")
	if env == "PRODUCTION" {
		return accounts.Prodconnection{Db: db}
	}

	return nil
}

func GetAssignmentsRepo() assignments.Repository {
	var env = os.Getenv("ENVIRONMENT")
	if env == "PRODUCTION" {
		return assignments.Prodconnection{Db: db}
	}

	return nil
}

func GetAccountAssignmentsRepo() intermediates.Repository {
	var env = os.Getenv("ENVIRONMENT")
	if env == "PRODUCTION" {
		return intermediates.Prodconnection{Db: db}
	}

	return nil
}

func GetSubmissionsRepo() submissions.Repository {
	var env = os.Getenv("ENVIRONMENT")
	if env == "PRODUCTION" {
		return submissions.Prodconnection{Db: db}
	}

	return nil
}

func createDB() {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Establishing contact with DMBS")
	statserr := utilities.StatsdClient.Inc("repo.base.createDB", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	env := os.Getenv("ENVIRONMENT")
	var pgr *gorm.DB
	var err error

	if env == "PRODUCTION" {
		pgr, err = gorm.Open(postgres.Open(string(os.Getenv("PSGR_CONNECTION"))), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	}

	if err != nil {
		slog.LogAttrs(context.Background(), constants.ErrorLevelFatal, "Unable to establish DBMS connection")
	}
	pgr.Exec("CREATE DATABASE " + os.Getenv("DB_NAME") + ";")
}

func setDBConnection() {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Establishing connection with DB")
	statserr := utilities.StatsdClient.Inc("repo.base.setDBConnection", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	env := os.Getenv("ENVIRONMENT")
	var pgrdb *gorm.DB
	var err error

	if env == "PRODUCTION" {
		pgrdb, err = gorm.Open(postgres.Open(string(os.Getenv("DB_CONNECTION"))), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	}

	if err != nil {
		slog.LogAttrs(context.Background(), constants.ErrorLevelFatal, "Unable to connect to "+os.Getenv("DB_NAME")+" DB")
	}

	db = pgrdb
}

func createModelsWithData() {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Creating schema and default accounts")
	statserr := utilities.StatsdClient.Inc("repo.base.createModelsWithData", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	env := os.Getenv("ENVIRONMENT")

	if env == "PRODUCTION" {
		err := db.AutoMigrate(&models.Account{})
		if err != nil {
			slog.LogAttrs(context.Background(), constants.ErrorLevelFatal, "Unable to create Accounts table")
		}

		err = db.AutoMigrate(&models.Assignment{})
		if err != nil {
			slog.LogAttrs(context.Background(), constants.ErrorLevelFatal, "Unable to create Assignments table")
		}

		err = db.AutoMigrate(&models.Submission{})
		if err != nil {
			slog.LogAttrs(context.Background(), constants.ErrorLevelFatal, "Unable to create Submissions table")
		}
	}

	lines := utilities.ReadCSV(os.Getenv("USERS_FILE"))
	lines = lines[1:]

	for _, line := range lines {
		cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))
		hpass, _ := bcrypt.GenerateFromPassword([]byte(line[3]), cost)
		pass := string(hpass)
		slog.LogAttrs(context.Background(), constants.ErrorLevelTrace, "Created password for "+line[2])
		uuid, _ := uuid2.NewUUID()
		usr := models.Account{ID: uuid.String(), First_Name: line[0], Last_Name: line[1], Email: line[2], Password: pass}
		psql := db.Create(&usr)
		if psql.Error != nil {
			slog.LogAttrs(context.Background(), slog.LevelError, "Error inserting user: "+usr.Email+" "+psql.Error.Error())
		}
	}
}

func Init() {
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		return
	}

	createDB()
	setDBConnection()
	createModelsWithData()
}

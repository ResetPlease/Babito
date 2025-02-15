package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/ResetPlease/Babito/internal/tools"
	_ "github.com/lib/pq"
)

type DatabaseCreds struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

func GetCredentials() DatabaseCreds {
	var creds DatabaseCreds
	creds.Host = tools.GetenvWithPanic("DATABASE_HOST")
	creds.Port = tools.GetenvWithPanic("DATABASE_PORT")
	creds.User = tools.GetenvWithPanic("DATABASE_USER")
	creds.Password = tools.GetenvWithPanic("DATABASE_PASSWORD")
	creds.DBname = tools.GetenvWithPanic("DATABASE_NAME")
	return creds
}

func databaseSetup(getCreds func() DatabaseCreds, logger *slog.Logger) (*sql.DB, error) {
	dbCreds := getCreds()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbCreds.Host, dbCreds.Port, dbCreds.User, dbCreds.Password, dbCreds.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("Failed connect to database", slog.Any("error", err))
		return nil, err
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(0)

	err = db.Ping()
	if err != nil {
		logger.Error("Failed healthcheck postgres", slog.Any("error", err))
		return nil, err
	}

	logger.Info("Success connection to database")

	return db, nil
}

type DatabaseController struct {
	DB     *sql.DB
	logger *slog.Logger
}

func (dc *DatabaseController) Close() {
	dc.DB.Close()
}

func NewDatabaseController(getCredentials func() DatabaseCreds, logger *slog.Logger) *DatabaseController {
	db, err := databaseSetup(getCredentials, logger)
	if err != nil {
		// because the database is a necessary element of this service
		panic(err)
	}
	return &DatabaseController{
		DB:     db,
		logger: logger,
	}
}

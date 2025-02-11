package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/ResetPlease/Babito/internal/tools"
	_ "github.com/lib/pq"
)

type DatabaseCreds struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func getCredentials() DatabaseCreds {
	var creds DatabaseCreds
	creds.host = tools.GetenvWithPanic("DATABASE_HOST")
	creds.port = tools.GetenvWithPanic("DATABASE_PORT")
	creds.user = tools.GetenvWithPanic("DATABASE_USER")
	creds.password = tools.GetenvWithPanic("DATABASE_PASSWORD")
	creds.dbname = tools.GetenvWithPanic("DATABASE_NAME")
	return creds
}

func databaseSetup(logger *slog.Logger) (*sql.DB, error) {
	dbCreds := getCredentials()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbCreds.host, dbCreds.port, dbCreds.user, dbCreds.password, dbCreds.dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("Failed connect to database", slog.Any("error", err))
		return nil, err
	}

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

func NewDatabaseController(logger *slog.Logger) *DatabaseController {
	db, err := databaseSetup(logger)
	if err != nil {
		// because the database is a necessary element of this service
		panic(err)
	}
	return &DatabaseController{
		DB:     db,
		logger: logger,
	}
}

package db

import (
	"database/sql"
	_ "embed"
	"errors"
	"log/slog"

	"github.com/ResetPlease/Babito/internal/models"
)

//go:embed queries/select_user_data_by_id.sql
var getUserDataByIDQuery string

//go:embed queries/select_user_data_by_username.sql
var getUserDataByUsernameQuery string

//go:embed queries/insert_new_user.sql
var insertNewUserDataQuery string

func (dc *DatabaseController) GetUserDataByUserID(id uint64) (*models.User, error) {
	stmt, err := dc.DB.Prepare(getUserDataByIDQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.HashedPassword, &user.Balance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		dc.logger.Error("Got error while getting user data", slog.Any("error", err))
		return nil, err
	}

	return &user, nil
}

func (dc *DatabaseController) GetUserDataByUsername(username string) (*models.User, error) {
	stmt, err := dc.DB.Prepare(getUserDataByUsernameQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.HashedPassword, &user.Balance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		dc.logger.Error("Got error while getting user data", slog.Any("error", err))
		return nil, err
	}

	return &user, nil
}

func (dc *DatabaseController) CreateNewUser(username string, hashedPassword string, balance int64) (*models.User, error) {
	var user models.User
	err := dc.DB.QueryRow(insertNewUserDataQuery, username, hashedPassword, balance).Scan(&user.ID, &user.Username, &user.HashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrDatabaseNotFound
		}
		dc.logger.Error("Got error while getting user data", slog.Any("error", err))
		return nil, err
	}

	return &user, nil
}

package db

import (
	_ "embed"
	"log/slog"

	"github.com/ResetPlease/Babito/internal/models"
)

//go:embed queries/select_user_data.sql
var getUserDataQuery string

func (dc *DatabaseController) GetUserDataByUserID(id uint64) (*models.User, error) {
	stmt, err := dc.DB.Prepare(getUserDataQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.HashedPassword, &user.Balance)

	if err != nil {
		dc.logger.Error("Got error while get user data", slog.Any("error", err))
		return nil, err
	}

	return &user, nil
}

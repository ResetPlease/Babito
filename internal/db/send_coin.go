package db

import (
	_ "embed"
	"log/slog"

	"github.com/ResetPlease/Babito/internal/models"
)

//go:embed queries/send_coin_transaction/lock_users.sql
var lockUsersQuery string

//go:embed queries/send_coin_transaction/send_coin.sql
var sendCoinQuery string

type LockedUser struct {
	ID      uint64
	Balance int64
}

func (dc *DatabaseController) SendCoinByUsername(fromUserID uint64, toUserUsername string, amount int64) (resErr error) {
	tx, err := dc.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if resErr != nil {
			rollbackError := tx.Rollback()
			if rollbackError != nil {
				dc.logger.Error("Got error while rollback", slog.Any("error", rollbackError))
			}
		}
	}()

	users := [2]LockedUser{}

	rows, err := tx.Query(lockUsersQuery, fromUserID, toUserUsername)
	if err != nil {
		return err
	}
	defer rows.Close()

	userCount := 0
	for rows.Next() {
		err = rows.Scan(&users[userCount].ID, &users[userCount].Balance)
		if err != nil {
			return err
		}
		userCount++
	}

	if err = rows.Err(); err != nil {
		return models.ErrUserNotFound
	}

	if userCount != 2 {
		return models.ErrUserNotFound
	}

	fromUser := LockedUser{}
	toUser := LockedUser{}
	if fromUserID == users[0].ID {
		fromUser = users[0]
		toUser = users[1]
	} else {
		fromUser = users[1]
		toUser = users[0]
	}

	if fromUser.Balance < amount {
		return models.ErrNotEnoughtFunds
	}

	_, err = tx.Exec(sendCoinQuery, fromUser.ID, toUser.ID, amount)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

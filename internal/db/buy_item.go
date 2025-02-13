package db

import (
	"database/sql"
	_ "embed"
	"errors"
	"log/slog"

	"github.com/ResetPlease/Babito/internal/models"
)

//go:embed queries/buy_item_transaction/lock_user.sql
var buyItemlockUserQuery string

//go:embed queries/buy_item_transaction/select_item_price_by_name.sql
var buyItemGetProductPrice string

//go:embed queries/buy_item_transaction/buy_item.sql
var buyItemUpdateUserBalance string

func (dc *DatabaseController) BuyItemByName(userID uint64, itemName string) (resErr error) {
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

	var itemPrice, userBalance int64
	err = tx.QueryRow(buyItemlockUserQuery, userID).Scan(&userBalance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrUserNotFound
		}
		return err
	}

	err = tx.QueryRow(buyItemGetProductPrice, itemName).Scan(&itemPrice)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrProductNotFound
		}
		return err
	}

	if itemPrice > userBalance {
		return models.ErrNotEnoughtFunds
	}

	_, err = tx.Exec(buyItemUpdateUserBalance, userID, itemPrice)
	if err != nil {
		return err
	}

	nullUserID := sql.NullInt64{Int64: 0, Valid: false}
	_, err = tx.Exec(createOperationQuery, userID, models.PURCHASE, itemPrice, nullUserID, itemName)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return nil
	}
	return nil
}

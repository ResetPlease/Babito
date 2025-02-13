package db

import (
	_ "embed"

	"github.com/ResetPlease/Babito/internal/models"
)

//go:embed queries/insert_new_operation.sql
var createOperationQuery string

//go:embed queries/select_transfer_history.sql
var getTransferHistoryQuery string

//go:embed queries/select_purchase_history.sql
var getPurchaseHistoryQuery string

func (dc *DatabaseController) GetTransfersByUserID(userID uint64) (models.Operations, error) {
	rows, err := dc.DB.Query(getTransferHistoryQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var operations models.Operations

	for rows.Next() {
		var operation models.Operation

		err := rows.Scan(&operation.Amount, &operation.TargetUsername, &operation.TargetUserID)
		if err != nil {
			return nil, err
		}

		operations = append(operations, operation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return operations, nil
}

func (dc *DatabaseController) GetPurchaseByUserID(userID uint64) (models.Operations, error) {
	rows, err := dc.DB.Query(getPurchaseHistoryQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var operations models.Operations

	for rows.Next() {
		var operation models.Operation

		err := rows.Scan(&operation.Amount, &operation.Item)
		if err != nil {
			return nil, err
		}

		operations = append(operations, operation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return operations, nil
}

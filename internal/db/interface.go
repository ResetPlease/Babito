package db

import "github.com/ResetPlease/Babito/internal/models"

type Database interface {
	GetUserDataByUserID(id uint64) (*models.User, error)
	GetUserDataByUsername(username string) (*models.User, error)
	CreateNewUser(username string, hashedPassword string, balance int64) (*models.User, error)

	SendCoinByUsername(fromUserID uint64, toUserUsername string, amount int64) error
	BuyItemByName(userID uint64, itemName string) error

	GetTransfersByUserID(userID uint64) (models.Operations, error)

	TestClearOperationHistory() error
	TestUpdateUsersBalance() error
}

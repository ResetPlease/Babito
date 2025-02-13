package models

import "errors"

var ErrDatabaseNotFound = errors.New("database not found")
var ErrUserNotFound = errors.New("user not found")
var ErrProductNotFound = errors.New("product not found")
var ErrNotEnoughtFunds = errors.New("not enough funds")

type OperationType string

const (
	TRANSFER OperationType = "transfer"
	PURCHASE OperationType = "purchase"
)

type Operation struct {
	UserID         uint64
	Type           OperationType
	Amount         int64
	TargetUserID   uint64
	TargetUsername string
	Item           string
}

type Operations []Operation

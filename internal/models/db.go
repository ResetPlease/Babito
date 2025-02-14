package models

import (
	"database/sql"
	"errors"
)

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
	UserID         uint64         `json:"user_id"`
	Username       string         `json:"username"`
	Type           OperationType  `json:"type"`
	Amount         int64          `json:"ampunt"`
	TargetUserID   sql.NullInt64  `json:"target_user_id"`
	TargetUsername sql.NullString `json:"target_username"`
	Item           sql.NullString `json:"item"`
}

type Operations []Operation

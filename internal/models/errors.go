package models

import "errors"

var ErrDatabaseNotFound = errors.New("database not found")
var ErrUserNotFound = errors.New("user not found")
var ErrNotEnoughtFunds = errors.New("not enough funds")

package db

import (
	"github.com/ResetPlease/Babito/internal/tools"
)

func (dc *DatabaseController) TestClearOperationHistory() error {
	_ = tools.GetenvWithPanic("TEST_MODE")

	_, err := dc.DB.Exec("DELETE FROM Operations;")
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseController) TestUpdateUsersBalance() error {
	_ = tools.GetenvWithPanic("TEST_MODE")

	_, err := dc.DB.Exec("UPDATE Users SET balance = $1;", 1000)
	if err != nil {
		return err
	}
	return nil
}

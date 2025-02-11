package models

type User struct {
	ID             uint64 `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Balance        int64  `json:"balance"`
}

func (u *User) CheckBalance() bool {
	return u.Balance >= 0
}

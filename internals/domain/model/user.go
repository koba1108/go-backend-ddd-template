package model

import (
	"errors"
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	IsWithdrew bool      `json:"isWithdrew"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (u *User) Withdraw() {
	u.IsWithdrew = true
	u.UpdatedAt = time.Now()
}

func (u *User) IsDeletable(ub *UserBank) error {
	// errorだけ返せば要件は満たすので、boolは不要
	if ub == nil {
		return errors.New("user bank is required")
	}
	if !u.IsWithdrew {
		return errors.New("user is not withdrew")
	}
	if u.ID != ub.UserId {
		return errors.New("user bank is must be same user id")
	}
	if ub.IsActive() {
		return errors.New("user bank is must be disabled")
	}
	return nil
}

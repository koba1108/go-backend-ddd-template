package model

import (
	"errors"
	"time"
)

type User struct {
	ID         int        `json:"id"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	IsWithdraw bool       `json:"isWithdraw"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

func NewUser(username, email string) (*User, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}
	if email == "" {
		return nil, errors.New("email is required")
	}
	return &User{
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (u *User) Withdraw() {
	u.IsWithdraw = true
	u.UpdatedAt = time.Now()
}

func (u *User) IsDeletable(ub *UserBank) error {
	// errorだけ返せば要件は満たすので、boolは不要
	if ub == nil {
		return errors.New("user bank is required")
	}
	if !u.IsWithdraw {
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

type UserCreateInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdateInput struct {
	ID         int     `json:"id"`
	Username   *string `json:"username"`
	Email      *string `json:"email"`
	IsWithdraw *bool   `json:"isWithdraw"`
}

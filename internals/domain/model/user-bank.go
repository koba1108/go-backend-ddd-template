package model

import "time"

type UserBankStatus string

const (
	UserBankStatusActive   UserBankStatus = "active"
	UserBankStatusDisabled UserBankStatus = "disabled"
)

type UserBank struct {
	ID        int            `json:"id"`
	UserId    int            `json:"userId"`
	BankName  string         `json:"bankName"`
	Status    UserBankStatus `json:"status"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

func (ub UserBank) IsActive() bool {
	return ub.Status == UserBankStatusActive
}

func (ub UserBank) Disable() {
	ub.Status = UserBankStatusDisabled
	ub.UpdatedAt = time.Now()
}

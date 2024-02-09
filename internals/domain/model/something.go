package model

import "time"

type Something struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

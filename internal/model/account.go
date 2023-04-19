package model

import "time"

type Account struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

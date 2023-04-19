package model

import "time"

type Transaction struct {
	ID        uint      `json:"id"`
	ItemID    uint      `json:"item_id"`
	UserID    uint      `json:"user_id"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

package model

import "time"

type ShareLink struct {
	ID        uint      `json:"id" `
	UserID    int       `json:"user_id"`
	ShareCode string    `json:"share_code"`
	FileIDs   string    `json:"file_ids"`
	Link      string    `json:"link"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

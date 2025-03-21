package model

import "time"

type ShareResource struct {
	ID           uint      `json:"id"`
	UserID       int       `json:"user_id"`
	ResourceType int       `json:"resource_type"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

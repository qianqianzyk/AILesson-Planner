package model

import (
	"gorm.io/gorm"
	"time"
)

type ConversationSession struct {
	ID        uint           `json:"id"`
	UserID    int            `json:"user_id"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ResponseConversationSession struct {
	ID                  int                   `json:"id"`
	UserID              int                   `json:"user_id"`
	Title               string                `json:"title"`
	CreatedAt           time.Time             `json:"created_at"`
	UpdatedAt           time.Time             `json:"updated_at"`
	DeletedAt           gorm.DeletedAt        `gorm:"index" json:"deleted_at"`
	ConversationMessage []ConversationMessage `json:"conversation_message"`
}

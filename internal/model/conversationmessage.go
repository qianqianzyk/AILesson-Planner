package model

import (
	"gorm.io/gorm"
	"time"
)

type ConversationMessage struct {
	ID        uint           `json:"id"`
	SessionID int            `json:"session_id"`
	UserID    int            `json:"user_id"`
	Role      string         `json:"role"` // "user" or "ai"
	Message   string         `json:"message"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Input struct {
	Messages []Message `json:"messages"`
}

type Parameters struct {
	ResultFormat string `json:"result_format"`
}

type RequestBody struct {
	Model      string     `json:"model"`
	Input      Input      `json:"input"`
	Parameters Parameters `json:"parameters"`
}

type Choice struct {
	FinishReason string `json:"finish_reason"`
	Message      struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
}

type Output struct {
	Text         string   `json:"text"`
	FinishReason string   `json:"finish_reason"`
	Choices      []Choice `json:"choices"`
}

type ApiResponse struct {
	StatusCode int    `json:"status_code"`
	RequestID  string `json:"request_id"`
	Output     Output `json:"output"`
	Usage      struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

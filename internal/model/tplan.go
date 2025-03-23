package model

import "time"

type TPlan struct {
	ID           uint      `json:"id"`
	UserID       int       `json:"user_id"`
	Subject      string    `json:"subject"`
	TextBookName string    `json:"textbook_name"`
	TopicHours   string    `json:"total_hours"`
	TopicName    string    `json:"topic_name"`
	TemplateFile string    `json:"template_file"`
	ResourceFile string    `json:"resource_file"`
	TextBookImg  string    `json:"textbook_img"`
	Description  string    `json:"description"`
	TPlanContent string    `json:"tplan_content"`
	TPlanUrl     string    `json:"tplan_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

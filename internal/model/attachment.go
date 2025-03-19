package model

import "time"

type Attachment struct {
	ID        uint      `json:"id"`
	UserID    int       `json:"user_id"`
	FileUrl   string    `json:"file_url"`
	FileName  string    `json:"file_name"`
	MD5       string    `json:"md5"`
	UploadID  string    `json:"upload_id"`
	CreatedAt time.Time `json:"created_at"`
}

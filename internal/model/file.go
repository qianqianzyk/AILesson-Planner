package model

import (
	"gorm.io/gorm"
	"time"
)

type File struct {
	ID        uint           `json:"id"`
	UserID    int            `json:"user_id"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Size      int            `json:"size"`
	FileType  string         `json:"file_type"`
	FType     int            `json:"f_type"`
	FileUrl   string         `json:"file_url"`
	IsDir     bool           `json:"is_dir"`
	IsCollect bool           `json:"is_collect"`
	ParentID  int            `json:"parent_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type DirectoryStats struct {
	FileCount      int64 `json:"file_count"`
	DirectoryCount int64 `json:"directory_count"`
	TotalSize      int64 `json:"total_size"`
}

const (
	VideoType    = 1
	AudioType    = 2
	DocumentType = 3
	ImageType    = 4
	OtherType    = 5
)

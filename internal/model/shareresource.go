package model

import "time"

type ShareResource struct {
	ID           uint      `json:"id"`
	UserID       int       `json:"user_id"`
	ResourceType int       `json:"resource_type"`
	ContentType  int       `json:"content_type"`
	CoverImg     string    `json:"cover_img"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	PdfUrl       string    `json:"pdf_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

const (
	SubjectModernChineseHistory = 1
	SubjectProgrammingC         = 2
	SubjectLinearAlgebraB       = 3
	SubjectCollegeEnglish       = 4
	SubjectNationalSecurity     = 5
	SubjectAdvancedMathematics  = 6
	SubjectCollegePhysics       = 7
)

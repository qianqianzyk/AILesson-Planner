package model

type Student struct {
	ID        uint   `json:"id"`
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	College   string `json:"college"`
	Class     string `json:"class"`
	Major     string `json:"major"`
}

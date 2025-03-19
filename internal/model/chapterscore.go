package model

type ChapterScore struct {
	ID        uint    `json:"id"`
	StudentID string  `json:"student_id"`
	CourseID  int     `json:"course_id"`
	Chapter   int     `json:"chapter"`
	Score     float64 `json:"score"`
}

type ChapterScoreWithAvg struct {
	Chapter      int     `json:"chapter"`
	StudentScore float64 `json:"student_score"`
	AvgScore     float64 `json:"avg_score"`
}

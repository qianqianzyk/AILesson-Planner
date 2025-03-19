package model

type Problem struct {
	ID          uint    `json:"id"`
	CourseID    int     `json:"course_id"`
	Chapter     int     `json:"chapter"`
	Tags        string  `json:"tags"`
	Question    string  `json:"question"`
	Options     string  `json:"options"`
	Answer      string  `json:"answer"`
	Score       float64 `json:"score"`
	ProblemType int     `json:"problem_type"`
}

const (
	SingleChoice   = 1
	MultipleChoice = 2
	TrueFalse      = 3
	FillBlank      = 4
	ShortAnswer    = 5
)

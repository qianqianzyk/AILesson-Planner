package model

type WrongProblem struct {
	ID           uint    `json:"id"`
	StudentID    int     `json:"student_id"`
	ProblemID    int     `json:"problem_id"`
	CreditEarned float64 `json:"credit_earned"`
	IsCorrect    bool    `json:"is_correct"`
}

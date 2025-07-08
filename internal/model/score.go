package model

type Score struct {
	ID           uint    `json:"id"`
	StudentID    string  `json:"student_id"`
	CourseID     int     `json:"course_id"`
	RegularScore float64 `json:"regular_score"`
	FinalScore   float64 `json:"final_score"`
	TotalScore   float64 `json:"total_score"`
	CreditEarned bool    `json:"credit_earned"`
	GradePoint   float64 `json:"grade_point"`
}

type StudentScore struct {
	StudentID    string  `json:"student_id"`
	Name         string  `json:"name"`
	Class        string  `json:"class"`
	Major        string  `json:"major"`
	College      string  `json:"college"`
	RegularScore float64 `json:"regular_score"`
	FinalScore   float64 `json:"final_score"`
	TotalScore   float64 `json:"total_score"`
	CreditEarned bool    `json:"credit_earned"`
	GradePoint   float64 `json:"grade_point"`
}

type CourseAvgScore struct {
	CourseID   uint    `json:"course_id"`
	CourseName string  `json:"course_name"`
	AvgScore   float64 `json:"avg_score"`
}

type ClassAvgGradePoint struct {
	AvgGradePoint float64 `json:"avg_grade_point"`
}

type CountPerformanceNumber struct {
	ExcellentCount int `json:"excellent_count"`
	GoodCount      int `json:"good_count"`
	PassCount      int `json:"pass_count"`
	FailCount      int `json:"fail_count"`
}

type TermGPA struct {
	AcademicYear string  `json:"academic_year"`
	AcademicTerm int     `json:"academic_term"`
	AvgGPA       float64 `json:"avg_gpa"`
	Rank         int     `json:"rank"`
	Percentile   float64 `json:"percentile"`
}

type StudentScores struct {
	StudentID    string  `json:"student_id"`
	Name         string  `json:"name"`
	Class        string  `json:"class"`
	Major        string  `json:"major"`
	College      string  `json:"college"`
	CourseID     int     `json:"course_id"`
	CourseName   string  `json:"course_name"`
	RegularScore float64 `json:"regular_score"`
	FinalScore   float64 `json:"final_score"`
	TotalScore   float64 `json:"total_score"`
	Credit       string  `json:"credit"`
	CreditEarned bool    `json:"credit_earned"`
	GradePoint   float64 `json:"grade_point"`
	AcademicYear string  `json:"academic_year"`
	AcademicTerm int     `json:"academic_term"`
}

type StudentTranscripts struct {
	StudentID string  `json:"student_id"`
	Name      string  `json:"name"`
	Class     string  `json:"class"`
	Ranking   int     `json:"ranking"`
	AvgGPA    float64 `json:"avg_gpa"`
}

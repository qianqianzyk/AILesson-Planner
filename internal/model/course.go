package model

type Course struct {
	ID                 uint   `json:"id"`
	UserID             int    `json:"user_id"`
	CourseNumber       string `json:"course_number"`
	CourseName         string `json:"course_name"`
	CourseImg          string `json:"course_img"`
	CourseIntroduction string `json:"course_introduction"`
	CourseClass        string `json:"course_class"`
	CourseType         int    `json:"course_type"`
	CourseAddr         string `json:"course_addr"`
	LecturerProfile    string `json:"lecturer_profile"`
	Credit             string `json:"credit"`
	AcademicYear       string `json:"academic_year"`
	AcademicTerm       int    `json:"academic_term"`
	Week               string `json:"week"`
	Weekday            string `json:"weekday"`
	Section            string `json:"section"`
	IsCompleted        bool   `json:"is_completed"`
}

const (
	CourseTypeRequired = 1 // 必修课
	CourseTypeElective = 2 // 选修课
	CourseTypeGeneral  = 3 // 任选课
	CourseTypePE       = 4 // 体育课
)

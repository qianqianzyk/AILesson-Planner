package dao

import (
	"context"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"gorm.io/gorm"
	"math"
	"strconv"
)

const (
	ExcellentThreshold = 90.0
	GoodThreshold      = 80.0
	PassThreshold      = 60.0
)

func (d *Dao) ImportStudentScores(ctx context.Context, rows [][]string, courseID,
	idxStudentID, idxName, idxCollege, idxClass, idxMajor, idxRegularScore, idxFinalScore, idxTotalScore int,
	hasName, hasCollege, hasClass, hasMajor, hasRegularScore, hasFinalScore, hasTotalScore bool) error {

	return d.orm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 收集所有studentID
		var studentIDs []string
		for _, row := range rows[1:] {
			if id := row[idxStudentID]; id != "" {
				studentIDs = append(studentIDs, id)
			}
		}
		// 预查询存在的学生
		var existingStudents []model.Student
		if err := tx.Where("student_id IN ?", studentIDs).Find(&existingStudents).Error; err != nil {
			return err
		}
		existingStudentMap := make(map[string]*model.Student)
		for i := range existingStudents {
			existingStudentMap[existingStudents[i].StudentID] = &existingStudents[i]
		}
		// 预查询存在的成绩
		var existingScores []model.Score
		if err := tx.Where("course_id = ? AND student_id IN ?", courseID, studentIDs).Find(&existingScores).Error; err != nil {
			return err
		}
		existingScoreMap := make(map[string]*model.Score)
		for i := range existingScores {
			existingScoreMap[existingScores[i].StudentID] = &existingScores[i]
		}
		// 处理学生数据
		var studentsToCreate []*model.Student
		var studentsToUpdate []*model.Student
		for _, row := range rows[1:] {
			studentID := row[idxStudentID]
			if studentID == "" {
				continue
			}

			if existingStudentMap[studentID] == nil {
				// 新增学生
				student := &model.Student{StudentID: studentID}
				if hasName {
					student.Name = row[idxName]
				}
				if hasCollege {
					student.College = row[idxCollege]
				}
				if hasClass {
					student.Class = row[idxClass]
				}
				if hasMajor {
					student.Major = row[idxMajor]
				}
				studentsToCreate = append(studentsToCreate, student)
			} else {
				// 更新学生
				student := &model.Student{StudentID: studentID}
				if hasName {
					student.Name = row[idxName]
				}
				if hasCollege {
					student.College = row[idxCollege]
				}
				if hasClass {
					student.Class = row[idxClass]
				}
				if hasMajor {
					student.Major = row[idxMajor]
				}
				studentsToUpdate = append(studentsToUpdate, student)
			}
		}
		// 批量创建学生
		if len(studentsToCreate) > 0 {
			if err := tx.Create(studentsToCreate).Error; err != nil {
				return err
			}
		}
		// 批量更新学生
		for _, student := range studentsToUpdate {
			updates := make(map[string]interface{})
			if hasName {
				updates["name"] = student.Name
			}
			if hasCollege {
				updates["college"] = student.College
			}
			if hasClass {
				updates["class"] = student.Class
			}
			if hasMajor {
				updates["major"] = student.Major
			}
			if len(updates) > 0 {
				if err := tx.Model(&model.Student{}).Where("student_id = ?", student.StudentID).Updates(updates).Error; err != nil {
					return err
				}
			}
		}

		// 处理成绩数据
		var scoresToCreate []*model.Score
		var scoresToUpdate []*model.Score
		for _, row := range rows[1:] {
			studentID := row[idxStudentID]
			if studentID == "" {
				continue
			}

			var score *model.Score
			if existingScore := existingScoreMap[studentID]; existingScore == nil {
				// 新增成绩
				score = &model.Score{StudentID: studentID, CourseID: courseID}
			} else {
				// 更新成绩（先复制原成绩的值）
				score = &model.Score{
					StudentID:    studentID,
					CourseID:     courseID,
					RegularScore: existingScore.RegularScore,
					FinalScore:   existingScore.FinalScore,
					TotalScore:   existingScore.TotalScore,
				}
			}

			// 更新需要修改的字段
			if hasRegularScore {
				score.RegularScore, _ = strconv.ParseFloat(row[idxRegularScore], 64)
			}
			if hasFinalScore {
				score.FinalScore, _ = strconv.ParseFloat(row[idxFinalScore], 64)
			}
			if hasTotalScore {
				score.TotalScore, _ = strconv.ParseFloat(row[idxTotalScore], 64)
			} else {
				score.TotalScore = math.Round(score.RegularScore*0.4 + score.FinalScore*0.6)
			}
			score.CreditEarned = score.TotalScore >= 60
			score.GradePoint = calculateGradePoint(score.TotalScore)

			if existingScoreMap[studentID] == nil {
				scoresToCreate = append(scoresToCreate, score)
			} else {
				scoresToUpdate = append(scoresToUpdate, score)
			}
		}

		// 批量创建成绩
		if len(scoresToCreate) > 0 {
			if err := tx.Create(scoresToCreate).Error; err != nil {
				return err
			}
		}
		// 批量更新成绩
		for _, score := range scoresToUpdate {
			updates := make(map[string]interface{})
			if hasRegularScore {
				updates["regular_score"] = score.RegularScore
			}
			if hasFinalScore {
				updates["final_score"] = score.FinalScore
			}
			updates["total_score"] = score.TotalScore
			updates["credit_earned"] = score.CreditEarned
			updates["grade_point"] = score.GradePoint

			if err := tx.Model(&model.Score{}).Where("student_id = ? AND course_id = ?", score.StudentID, score.CourseID).Updates(updates).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (d *Dao) GetStudentScoresByClass(ctx context.Context, courseID int, class string) ([]model.StudentScore, error) {
	var results []model.StudentScore
	err := d.orm.WithContext(ctx).
		Table("scores").
		Select("students.student_id, students.name, students.class, students.major, students.college, scores.regular_score, scores.final_score, scores.total_score, scores.credit_earned, scores.grade_point").
		Joins("JOIN students ON scores.student_id = students.student_id").
		Where("scores.course_id = ? AND students.class = ?", courseID, class).
		Scan(&results).Error
	return results, err
}

func (d *Dao) GetStudentScoresByCourse(ctx context.Context, courseID int) ([]model.StudentScore, error) {
	var results []model.StudentScore
	err := d.orm.WithContext(ctx).
		Table("scores").
		Select("students.student_id, students.name, students.class, students.major, students.college, scores.regular_score, scores.final_score, scores.total_score, scores.credit_earned, scores.grade_point").
		Joins("JOIN students ON scores.student_id = students.student_id").
		Where("scores.course_id = ?", courseID).
		Scan(&results).Error
	return results, err
}

func (d *Dao) BatchUpdateStudentScores(ctx context.Context, data []model.StudentScore, courseID int) error {
	tx := d.orm.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, entry := range data {
		err := tx.Model(&model.Student{}).
			Where("student_id = ?", entry.StudentID).
			Updates(map[string]interface{}{
				"name":    entry.Name,
				"class":   entry.Class,
				"major":   entry.Major,
				"college": entry.College,
			}).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("更新学生信息失败: %w", err)
		}

		var totalScore float64
		if entry.TotalScore == 0 {
			totalScore = math.Round(entry.RegularScore*0.4 + entry.FinalScore*0.6)
		} else {
			totalScore = entry.TotalScore
		}
		creditEarned := totalScore >= 60
		gradePoint := calculateGradePoint(totalScore)
		err = tx.Model(&model.Score{}).
			Where("student_id = ? AND course_id = ?", entry.StudentID, courseID).
			Updates(map[string]interface{}{
				"regular_score": entry.RegularScore,
				"final_score":   entry.FinalScore,
				"total_score":   totalScore,
				"credit_earned": creditEarned,
				"grade_point":   gradePoint,
			}).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("更新学生成绩失败: %w", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	return nil
}

func calculateGradePoint(totalScore float64) float64 {
	if totalScore >= 90 {
		return 4.0 + (totalScore-90)/10
	} else if totalScore >= 80 {
		return 3.0 + (totalScore-80)/10
	} else if totalScore >= 70 {
		return 2.0 + (totalScore-70)/10
	} else if totalScore >= 60 {
		return 1.0 + (totalScore-60)/10
	} else {
		return 0.0
	}
}

func (d *Dao) BatchDeleteScores(ctx context.Context, studentIDs []string, courseID int) error {
	if len(studentIDs) == 0 {
		return nil
	}

	err := d.orm.WithContext(ctx).
		Where("student_id IN (?) AND course_id = ?", studentIDs, courseID).
		Delete(&model.Score{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) BatchUpsertStudentScores(ctx context.Context, data []model.Score, courseID int) error {
	if len(data) == 0 {
		return nil
	}

	tx := d.orm.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, entry := range data {
		if entry.TotalScore == 0 {
			entry.TotalScore = math.Round(entry.RegularScore*0.4 + entry.FinalScore*0.6)
		}
		entry.CreditEarned = entry.TotalScore >= 60
		entry.GradePoint = calculateGradePoint(entry.TotalScore)

		var existing model.Score
		err := tx.Where("student_id = ? AND course_id = ?", entry.StudentID, courseID).First(&existing).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			tx.Rollback()
			return fmt.Errorf("查询成绩记录失败: %w", err)
		}

		if err == gorm.ErrRecordNotFound {
			if err = tx.Create(&entry).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("成绩记录创建失败: %w", err)
			}
		} else {
			if err = tx.Model(&existing).Updates(map[string]interface{}{
				"regular_score": entry.RegularScore,
				"final_score":   entry.FinalScore,
				"total_score":   entry.TotalScore,
				"credit_earned": entry.CreditEarned,
				"grade_point":   entry.GradePoint,
			}).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("成绩记录更新失败: %w", err)
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	return nil
}

func (d *Dao) DeleteScoreByID(ctx context.Context, courseID int) error {
	result := d.orm.WithContext(ctx).
		Where("course_id = ?", courseID).
		Delete(&model.Score{})
	return result.Error
}

func (d *Dao) GetClassPerformance(ctx context.Context, academicYear string, academicTerm int, courseClass string) ([]model.CourseAvgScore, model.ClassAvgGradePoint, error) {
	var courseAvgScores []model.CourseAvgScore
	var classAvgGradePoint model.ClassAvgGradePoint

	// 计算每门课程的平均分
	err := d.orm.WithContext(ctx).
		Table("scores").
		Select("courses.id as course_id, courses.course_name, AVG(scores.total_score) as avg_score").
		Joins("JOIN courses ON scores.course_id = courses.id").
		Joins("JOIN students ON scores.student_id = students.student_id").
		Where("courses.academic_year = ? AND courses.academic_term = ? AND students.class = ?", academicYear, academicTerm, courseClass).
		Group("courses.id, courses.course_name").
		Scan(&courseAvgScores).Error
	if err != nil {
		return nil, model.ClassAvgGradePoint{}, fmt.Errorf("查询课程平均分失败: %w", err)
	}

	// 计算班级的平均绩点
	err = d.orm.WithContext(ctx).
		Table("scores").
		Select("AVG(scores.grade_point) as avg_grade_point").
		Joins("JOIN courses ON scores.course_id = courses.id").
		Joins("JOIN students ON scores.student_id = students.student_id").
		Where("courses.academic_year = ? AND courses.academic_term = ? AND students.class = ?", academicYear, academicTerm, courseClass).
		Scan(&classAvgGradePoint).Error
	if err != nil {
		return nil, model.ClassAvgGradePoint{}, fmt.Errorf("查询班级平均绩点失败: %w", err)
	}

	return courseAvgScores, classAvgGradePoint, nil
}

func (d *Dao) GetScoreOverview(ctx context.Context, academicYear string, academicTerm int, className string, courseName string) (model.CountPerformanceNumber, error) {
	var courseID int

	if courseName != "" {
		err := d.orm.WithContext(ctx).
			Model(model.Course{}).
			Select("id").
			Where("academic_year = ? AND academic_term = ? AND course_name = ?", academicYear, academicTerm, courseName).
			First(&courseID).Error
		if err != nil {
			return model.CountPerformanceNumber{}, fmt.Errorf("查询课程ID失败: %w", err)
		}
	} else {
		err := d.orm.WithContext(ctx).
			Model(model.Course{}).
			Select("id").
			Where("academic_year = ? AND academic_term = ?", academicYear, academicTerm).
			First(&courseID).Error
		if err != nil {
			return model.CountPerformanceNumber{}, fmt.Errorf("查询课程ID失败: %w", err)
		}
	}

	var overview model.CountPerformanceNumber
	err := d.orm.WithContext(ctx).
		Table("scores").
		Select("SUM(CASE WHEN total_score >= ? THEN 1 ELSE 0 END) AS excellent_count, "+
			"SUM(CASE WHEN total_score >= ? AND total_score < ? THEN 1 ELSE 0 END) AS good_count, "+
			"SUM(CASE WHEN total_score >= ? AND total_score < ? THEN 1 ELSE 0 END) AS pass_count, "+
			"SUM(CASE WHEN total_score < ? THEN 1 ELSE 0 END) AS fail_count",
			ExcellentThreshold, GoodThreshold, ExcellentThreshold, PassThreshold, GoodThreshold, PassThreshold).
		Joins("JOIN students ON scores.student_id = students.student_id").
		Where("scores.course_id = ? AND students.class = ?", courseID, className).
		Scan(&overview).Error
	if err != nil {
		return model.CountPerformanceNumber{}, fmt.Errorf("查询成绩概况失败: %w", err)
	}

	return overview, nil
}

func (d *Dao) GetClassListByUserID(ctx context.Context, userID int) ([]string, error) {
	var courseIDs []int
	err := d.orm.WithContext(ctx).
		Table("courses").
		Select("id").
		Where("user_id = ?", userID).
		Pluck("id", &courseIDs).Error
	if err != nil {
		return nil, fmt.Errorf("查询课程列表失败: %w", err)
	}
	if len(courseIDs) == 0 {
		return []string{}, nil
	}

	var studentIDs []string
	err = d.orm.WithContext(ctx).
		Table("scores").
		Select("DISTINCT student_id").
		Where("course_id IN ?", courseIDs).
		Pluck("student_id", &studentIDs).Error
	if err != nil {
		return nil, fmt.Errorf("查询学生ID失败: %w", err)
	}
	if len(studentIDs) == 0 {
		return []string{}, nil
	}

	var classList []string
	err = d.orm.WithContext(ctx).
		Table("students").
		Select("DISTINCT class").
		Where("student_id IN ?", studentIDs).
		Pluck("class", &classList).Error
	if err != nil {
		return nil, fmt.Errorf("查询班级列表失败: %w", err)
	}

	return classList, nil
}

func (d *Dao) GetClassListByCourseID(ctx context.Context, courseID int) ([]string, error) {
	var studentIDs []string
	err := d.orm.WithContext(ctx).
		Table("scores").
		Select("DISTINCT student_id").
		Where("course_id = ?", courseID).
		Pluck("student_id", &studentIDs).Error
	if err != nil {
		return nil, fmt.Errorf("查询学生ID失败: %w", err)
	}
	if len(studentIDs) == 0 {
		return []string{}, nil
	}

	var classList []string
	err = d.orm.WithContext(ctx).
		Table("students").
		Select("DISTINCT class").
		Where("student_id IN ?", studentIDs).
		Pluck("class", &classList).Error
	if err != nil {
		return nil, fmt.Errorf("查询班级列表失败: %w", err)
	}

	return classList, nil
}

func (d *Dao) GetClassList(ctx context.Context) ([]string, error) {
	var classList []string

	err := d.orm.WithContext(ctx).
		Table("students").
		Select("DISTINCT class").
		Pluck("class", &classList).Error
	if err != nil {
		return nil, fmt.Errorf("查询班级列表失败: %w", err)
	}

	return classList, nil
}

func (d *Dao) GetStudentsByUserAndClass(ctx context.Context, userID int, class string) ([]model.Student, error) {
	var courseIDs []int
	err := d.orm.WithContext(ctx).
		Table("courses").
		Select("id").
		Where("user_id = ?", userID).
		Pluck("id", &courseIDs).Error
	if err != nil {
		return nil, fmt.Errorf("查询课程ID失败: %w", err)
	}
	if len(courseIDs) == 0 {
		return nil, nil
	}

	var studentIDs []string
	err = d.orm.WithContext(ctx).
		Table("scores").
		Select("DISTINCT student_id").
		Where("course_id IN ?", courseIDs).
		Pluck("student_id", &studentIDs).Error
	if err != nil {
		return nil, fmt.Errorf("查询学生ID失败: %w", err)
	}
	if len(studentIDs) == 0 {
		return nil, nil
	}

	var students []model.Student
	err = d.orm.WithContext(ctx).
		Table("students").
		Where("student_id IN ? AND class = ?", studentIDs, class).
		Find(&students).Error
	if err != nil {
		return nil, fmt.Errorf("查询学生姓名失败: %w", err)
	}

	return students, nil
}

func (d *Dao) GetStudentGPAAndRank(ctx context.Context, studentID string) ([]model.TermGPA, error) {
	var gpaResults []model.TermGPA
	// 查询该学生各学期的平均绩点
	err := d.orm.WithContext(ctx).
		Table("scores").
		Select("courses.academic_year, courses.academic_term, AVG(scores.grade_point) as avg_gpa").
		Joins("JOIN courses ON scores.course_id = courses.id").
		Where("scores.student_id = ?", studentID).
		Group("courses.academic_year, courses.academic_term").
		Scan(&gpaResults).Error
	if err != nil {
		return nil, fmt.Errorf("查询学生绩点失败: %w", err)
	}
	// 查询该学生的年级、专业信息
	var studentInfo struct {
		Class string `json:"class"`
		Major string `json:"major"`
	}
	err = d.orm.WithContext(ctx).
		Table("students").
		Select("class, major").
		Where("student_id = ?", studentID).
		First(&studentInfo).Error
	if err != nil {
		return nil, fmt.Errorf("查询学生信息失败: %w", err)
	}
	// 遍历每个学期的绩点，计算排名
	for i, term := range gpaResults {
		// 计算年级排名
		var rank int
		err = d.orm.WithContext(ctx).
			Raw(`
				SELECT COUNT(*) + 1 FROM (
					SELECT scores.student_id, AVG(scores.grade_point) as avg_gpa
					FROM scores
					JOIN courses ON scores.course_id = courses.id
					JOIN students ON scores.student_id = students.student_id
					WHERE students.class = ? AND courses.academic_year = ? AND courses.academic_term = ?
					GROUP BY scores.student_id
				) as grade_gpa
				WHERE grade_gpa.avg_gpa > ?
			`, studentInfo.Class, term.AcademicYear, term.AcademicTerm, term.AvgGPA).
			Scan(&rank).Error
		if err != nil {
			return nil, fmt.Errorf("查询年级排名失败: %w", err)
		}
		gpaResults[i].Rank = rank
		// 计算专业前百分比
		var totalInMajor int
		err = d.orm.WithContext(ctx).
			Raw(`
				SELECT COUNT(*) FROM (
					SELECT student_id
					FROM students
					WHERE major = ?
				) as major_students
			`, studentInfo.Major).
			Scan(&totalInMajor).Error
		if err != nil {
			return nil, fmt.Errorf("查询专业总人数失败: %w", err)
		}

		if totalInMajor > 0 {
			gpaResults[i].Percentile = float64(rank) / float64(totalInMajor) * 100
		} else {
			gpaResults[i].Percentile = 0
		}
	}

	return gpaResults, nil
}

func (d *Dao) GetStudentTranscripts(ctx context.Context, academicYear, name, studentID string, academicTerm int) ([]model.StudentTranscripts, error) {
	var transcripts []model.StudentTranscripts

	err := d.orm.WithContext(ctx).
		Table("students AS s").
		Select("s.student_id, s.name, s.class, COALESCE(AVG(sc.grade_point), 0) AS avg_gpa").
		Joins("LEFT JOIN scores AS sc ON s.student_id = sc.student_id").
		Joins("LEFT JOIN courses AS c ON sc.course_id = c.id").
		Where("c.academic_year = ? AND c.academic_term = ?", academicYear, academicTerm).
		Where("s.name LIKE ? AND s.student_id LIKE ?", "%"+name+"%", "%"+studentID+"%").
		Group("s.student_id, s.name, s.class").
		Order("avg_gpa DESC").
		Scan(&transcripts).Error
	if err != nil {
		return nil, fmt.Errorf("查询成绩单失败: %w", err)
	}

	for i := range transcripts {
		transcripts[i].Ranking = i + 1
	}

	return transcripts, nil
}

func (d *Dao) GetStudentCourses(ctx context.Context, studentID string) ([]model.StudentScores, error) {
	var studentScores []model.StudentScores
	err := d.orm.WithContext(ctx).
		Table("scores").
		Select("students.student_id, students.name, students.class, students.major, students.college, scores.course_id, courses.course_name, scores.regular_score, scores.final_score, scores.total_score, courses.credit, scores.credit_earned, scores.grade_point, courses.academic_year, courses.academic_term").
		Joins("JOIN students ON scores.student_id = students.student_id").
		Joins("JOIN courses ON scores.course_id = courses.id").
		Where("scores.student_id = ?", studentID).
		Order("CAST(SUBSTRING(courses.academic_year, 1, 4) AS SIGNED) DESC, courses.academic_term DESC").
		Scan(&studentScores).Error

	if err != nil {
		return nil, fmt.Errorf("查询学生课程记录失败: %w", err)
	}

	return studentScores, nil
}

func (d *Dao) GetChapterScoresWithAvg(ctx context.Context, studentID string, courseID int) ([]model.ChapterScoreWithAvg, error) {
	var results []model.ChapterScoreWithAvg
	// 查询该学生在该课程每一章的成绩
	err := d.orm.WithContext(ctx).
		Table("chapter_scores").
		Select("chapter, score AS student_score").
		Where("student_id = ? AND course_id = ?", studentID, courseID).
		Scan(&results).Error
	if err != nil {
		return nil, fmt.Errorf("查询学生的章节成绩失败: %w", err)
	}
	// 查询该课程每一章的班级平均成绩
	var avgScores []struct {
		Chapter  int     `json:"chapter"`
		Class    string  `json:"class"`
		AvgScore float64 `json:"avg_score"`
	}
	err = d.orm.WithContext(ctx).
		Table("chapter_scores").
		Select("chapter_scores.chapter, students.class, AVG(chapter_scores.score) AS avg_score").
		Joins("JOIN students ON chapter_scores.student_id = students.student_id").
		Where("chapter_scores.course_id = ?", courseID).
		Group("chapter_scores.chapter, students.class").
		Scan(&avgScores).Error
	if err != nil {
		return nil, fmt.Errorf("查询课程每章的班级平均成绩失败: %w", err)
	}
	// 将班级平均成绩添加到每一章的成绩中
	for i := range results {
		for _, avg := range avgScores {
			if results[i].Chapter == avg.Chapter {
				results[i].AvgScore = avg.AvgScore
				break
			}
		}
	}
	return results, nil
}

func (d *Dao) GetStudentsByClasses(ctx context.Context, classes []string) ([]model.Student, error) {
	var students []model.Student
	err := d.orm.WithContext(ctx).Where("class IN ?", classes).Find(&students).Error
	return students, err
}

func (d *Dao) CreateScoresBatch(ctx context.Context, scores []*model.Score) error {
	if err := d.orm.WithContext(ctx).Create(&scores).Error; err != nil {
		return fmt.Errorf("批量插入学生创建记录失败: %w", err)
	}
	return nil
}

func (d *Dao) CreateScore(ctx context.Context, score *model.Score) error {
	if err := d.orm.WithContext(ctx).Create(&score).Error; err != nil {
		return err
	}
	return nil
}

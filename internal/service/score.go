package service

import (
	"bytes"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func ImportStudentScores(rows [][]string, courseID,
	idxStudentID, idxName, idxCollege, idxClass, idxMajor, idxRegularScore, idxFinalScore, idxTotalScore int,
	hasName, hasCollege, hasClass, hasMajor, hasRegularScore, hasFinalScore, hasTotalScore bool) error {
	err := d.ImportStudentScores(ctx, rows, courseID,
		idxStudentID, idxName, idxCollege, idxClass, idxMajor, idxRegularScore, idxFinalScore, idxTotalScore,
		hasName, hasCollege, hasClass, hasMajor, hasRegularScore, hasFinalScore, hasTotalScore)
	return err
}

func GetStudentScoreByClass(courseID int, class string) ([]model.StudentScore, error) {
	scores, err := d.GetStudentScoresByClass(ctx, courseID, class)
	return scores, err
}

func GetStudentScore(courseID int) ([]model.StudentScore, error) {
	scores, err := d.GetStudentScoresByCourse(ctx, courseID)
	return scores, err
}

func ExportStudentScores(courseClass string, scores []model.StudentScore) (string, error) {
	// 创建一个新的Excel文件
	f := excelize.NewFile()
	sheetName := "Sheet1"
	// 设置标题样式（加粗 + 居中）
	titleStyleID, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Size: 14},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return "", errors.New("设置标题样式失败: " + err.Error())
	}
	// 设置表头样式（加粗）
	headerStyleID, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
	})
	if err != nil {
		return "", errors.New("设置表头样式失败: " + err.Error())
	}
	// 合并单元格 A1 到 I1
	if err := f.MergeCell(sheetName, "A1", "J1"); err != nil {
		return "", errors.New("合并单元格失败: " + err.Error())
	}
	// 设置第一行标题
	f.SetCellValue(sheetName, "A1", courseClass+" 成绩")
	f.SetCellStyle(sheetName, "A1", "I1", titleStyleID)
	// 设置第二行表头
	headers := []string{"序号", "姓名", "学号", "班级", "专业", "学院", "平时成绩", "期末成绩", "最终成绩", "获得学分", "绩点"}
	for i, header := range headers {
		colName, _ := excelize.ColumnNumberToName(i + 1)
		cell := colName + "2"
		f.SetCellValue(sheetName, cell, header)
		f.SetCellStyle(sheetName, cell, cell, headerStyleID)
	}
	// 填充数据（从第3行开始）
	for i, score := range scores {
		rowIndex := strconv.Itoa(i + 3)
		data := []any{
			i + 1,
			score.Name,
			score.StudentID,
			score.Class,
			score.Major,
			score.College,
			score.RegularScore,
			score.FinalScore,
			score.TotalScore,
			map[bool]string{true: "是", false: "否"}[score.CreditEarned],
			score.GradePoint,
		}
		for j, value := range data {
			colName, _ := excelize.ColumnNumberToName(j + 1)
			f.SetCellValue(sheetName, colName+rowIndex, value)
		}
	}
	// 将 Excel 文件写入内存
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return "", errors.New("写入内存失败: " + err.Error())
	}
	// 生成 MinIO 存储路径
	objectKey := "scores/" + courseClass + "_成绩.xlsx"
	// 调用 MinIO 上传
	url, err := d.PutObject(objectKey, &buf, int64(buf.Len()), "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	if err != nil {
		return "", errors.New("上传到 MinIO 失败: " + err.Error())
	}
	return url, nil
}

func UpdateStudentScores(data []model.StudentScore, courseID int) error {
	err := d.BatchUpdateStudentScores(ctx, data, courseID)
	return err
}

func DeleteStudentScores(studentIDs []string, courseID int) error {
	err := d.BatchDeleteScores(ctx, studentIDs, courseID)
	return err
}

func UpsertStudentScore(data []model.Score, courseID int) error {
	err := d.BatchUpsertStudentScores(ctx, data, courseID)
	return err
}

func GetClassPerformance(academicYear string, academicTerm int, courseClass string) ([]model.CourseAvgScore, model.ClassAvgGradePoint, error) {
	courseAvgScores, classAvgGradePoint, err := d.GetClassPerformance(ctx, academicYear, academicTerm, courseClass)
	return courseAvgScores, classAvgGradePoint, err
}

func GetScoreOverview(academicYear string, academicTerm int, className string, courseName string) (model.CountPerformanceNumber, error) {
	overview, err := d.GetScoreOverview(ctx, academicYear, academicTerm, className, courseName)
	return overview, err
}

func GetClassListByUserID(userID int) ([]string, error) {
	classList, err := d.GetClassListByUserID(ctx, userID)
	return classList, err
}

func GetClassListByCourseID(courseID int) ([]string, error) {
	classList, err := d.GetClassListByCourseID(ctx, courseID)
	return classList, err
}

func GetClassList() ([]string, error) {
	classList, err := d.GetClassList(ctx)
	return classList, err
}

func GetStudentsByUserAndClass(userID int, class string) ([]model.Student, error) {
	studentList, err := d.GetStudentsByUserAndClass(ctx, userID, class)
	return studentList, err
}

func GetStudentGPAAndRank(studentID string) ([]model.TermGPA, error) {
	termGPAs, err := d.GetStudentGPAAndRank(ctx, studentID)
	return termGPAs, err
}

func GetStudentCourses(studentID string) ([]model.StudentScores, error) {
	studentScores, err := d.GetStudentCourses(ctx, studentID)
	return studentScores, err
}

func GetChapterScoresWithAvg(studentID string, courseID int) ([]model.ChapterScoreWithAvg, error) {
	result, err := d.GetChapterScoresWithAvg(ctx, studentID, courseID)
	return result, err
}

func CreateScoresBatch(courseID int, classes []string) error {
	students, err := d.GetStudentsByClasses(ctx, classes)
	if err != nil {
		return err
	}

	var scores []*model.Score
	for _, student := range students {
		scores = append(scores, &model.Score{
			StudentID:    student.StudentID,
			CourseID:     courseID,
			RegularScore: 0,
			FinalScore:   0,
			TotalScore:   0,
			CreditEarned: false,
			GradePoint:   0,
		})
	}

	err = d.CreateScoresBatch(ctx, scores)
	if err != nil {
		return err
	}
	return nil
}

func CreateScore(score *model.Score) error {
	err := d.CreateScore(ctx, score)
	return err
}

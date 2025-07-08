package dao

import (
	"context"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
)

func (d *Dao) CreateStudent(ctx context.Context, student *model.Student) error {
	result := d.orm.WithContext(ctx).Model(&model.Student{}).Create(student)
	return result.Error
}

func (d *Dao) UpdateStudent(ctx context.Context, student *model.Student) error {
	result := d.orm.WithContext(ctx).Save(&student)
	return result.Error
}

func (d *Dao) GetStudentByStudentID(ctx context.Context, studentID string) (*model.Student, error) {
	var student model.Student
	result := d.orm.WithContext(ctx).Model(&model.Student{}).Where("student_id = ?", studentID).First(&student)
	return &student, result.Error
}

func (d *Dao) GetStudentByStudentIDAndCourseID(ctx context.Context, studentID string, courseID int) (*model.Score, error) {
	var score model.Score
	result := d.orm.WithContext(ctx).Model(&model.Score{}).Where("student_id = ? AND course_id = ?", studentID, courseID).First(&score)
	return &score, result.Error
}

func (d *Dao) GetStudentsByCourse(ctx context.Context, courseID int, class string) ([]model.Student, error) {
	var studentIDs []string

	err := d.orm.WithContext(ctx).
		Table("scores").
		Select("DISTINCT student_id").
		Where("course_id = ?", courseID).
		Pluck("student_id", &studentIDs).Error
	if err != nil {
		return nil, fmt.Errorf("查询学生ID失败: %w", err)
	}

	fmt.Println(studentIDs)

	if len(studentIDs) == 0 {
		return []model.Student{}, nil
	}

	var students []model.Student
	err = d.orm.WithContext(ctx).
		Table("students").
		Where("student_id IN ? AND class = ?", studentIDs, class).
		Find(&students).Error
	if err != nil {
		return nil, fmt.Errorf("查询学生信息失败: %w", err)
	}

	return students, nil
}

func (d *Dao) DeleteStudents(ctx context.Context, studentIDs []string, courseID int) error {
	if len(studentIDs) == 0 {
		return nil
	}

	err := d.orm.WithContext(ctx).
		Where("student_id IN (?) AND course_id = ?", studentIDs, courseID).
		Delete(&model.Score{}).Error
	if err != nil {
		return err
	}

	var remainingStudentIDs []string
	err = d.orm.WithContext(ctx).
		Model(&model.Score{}).
		Where("student_id IN (?)", studentIDs).
		Pluck("student_id", &remainingStudentIDs).Error
	if err != nil {
		return err
	}

	studentIDSet := make(map[string]struct{}, len(studentIDs))
	for _, id := range studentIDs {
		studentIDSet[id] = struct{}{}
	}

	for _, id := range remainingStudentIDs {
		delete(studentIDSet, id)
	}

	finalStudentIDs := make([]string, 0, len(studentIDSet))
	for id := range studentIDSet {
		finalStudentIDs = append(finalStudentIDs, id)
	}

	if len(finalStudentIDs) == 0 {
		return nil
	}

	err = d.orm.WithContext(ctx).
		Where("student_id IN (?)", finalStudentIDs).
		Delete(&model.Student{}).Error
	if err != nil {
		return err
	}

	return nil
}

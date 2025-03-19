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

func (d *Dao) DeleteStudents(ctx context.Context, studentIDs []string) error {
	if len(studentIDs) == 0 {
		return nil
	}

	err := d.orm.WithContext(ctx).
		Where("student_id IN (?)", studentIDs).
		Delete(&model.Student{}).Error
	if err != nil {
		return err
	}

	err = d.orm.WithContext(ctx).
		Where("student_id IN (?)", studentIDs).
		Delete(&model.Score{}).Error
	if err != nil {
		return err
	}

	return nil
}

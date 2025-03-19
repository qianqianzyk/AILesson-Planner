package service

import "github.com/qianqianzyk/AILesson-Planner/internal/model"

func CreateStudent(student *model.Student) error {
	err := d.CreateStudent(ctx, student)
	return err
}

func GetStudentByStudentID(studentID string) (*model.Student, error) {
	student, err := d.GetStudentByStudentID(ctx, studentID)
	return student, err
}

func UpdateStudent(student *model.Student) error {
	err := d.UpdateStudent(ctx, student)
	return err
}

func GetStudentsByCourse(courseID int, class string) ([]model.Student, error) {
	students, err := d.GetStudentsByCourse(ctx, courseID, class)
	return students, err
}

func DeleteStudents(studentIDs []string) error {
	err := d.DeleteStudents(ctx, studentIDs)
	return err
}

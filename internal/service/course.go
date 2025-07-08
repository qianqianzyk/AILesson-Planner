package service

import "github.com/qianqianzyk/AILesson-Planner/internal/model"

func CreateCourse(course *model.Course) error {
	err := d.CreateCourse(ctx, course)
	return err
}

func UpdateCourse(course *model.Course) error {
	err := d.UpdateCourse(ctx, course)
	return err
}

func GetCourseList(userID int) ([]model.Course, error) {
	courses, err := d.GetCourseList(ctx, userID)
	return courses, err
}

func GetCourseByID(id int) (*model.Course, error) {
	course, err := d.GetCourseByID(ctx, id)
	return course, err
}

func DeleteCourseByID(id, userID int) error {
	err := d.DeleteCourseByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}
	err = d.DeleteScoreByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func GetCoursesByWeek(academicYear string, academicTerm, userID, week int) ([]model.Course, error) {
	courses, err := d.GetCoursesByWeek(ctx, academicYear, academicTerm, userID, week)
	return courses, err
}

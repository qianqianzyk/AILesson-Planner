package dao

import (
	"context"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"sort"
	"strconv"
	"strings"
)

func (d *Dao) CreateCourse(ctx context.Context, course *model.Course) error {
	result := d.orm.WithContext(ctx).Model(&model.Course{}).Create(course)
	return result.Error
}

func (d *Dao) GetCourseList(ctx context.Context, userID int) ([]model.Course, error) {
	var courses []model.Course
	err := d.orm.WithContext(ctx).Model(&model.Course{}).
		Where("user_id = ?", userID).
		Order("is_completed ASC").Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (d *Dao) GetCourseByID(ctx context.Context, id int) (*model.Course, error) {
	var course *model.Course
	result := d.orm.WithContext(ctx).Model(&model.Course{}).Where("id = ?", id).First(&course)
	if result.Error != nil {
		return nil, result.Error
	}
	return course, nil
}

func (d *Dao) DeleteCourseByID(ctx context.Context, id int) error {
	result := d.orm.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Course{})
	return result.Error
}

func (d *Dao) DeleteCourseByIDAndUserID(ctx context.Context, id, userID int) error {
	result := d.orm.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&model.Course{})
	return result.Error
}

func (d *Dao) UpdateCourse(ctx context.Context, course *model.Course) error {
	result := d.orm.WithContext(ctx).Save(&course)
	return result.Error
}

func (d *Dao) GetCoursesByWeek(ctx context.Context, academicYear string, academicTerm int, userID int, week int) ([]model.Course, error) {
	var courses []model.Course

	err := d.orm.WithContext(ctx).
		Table("courses").
		Where("academic_year = ? AND academic_term = ? AND user_id = ?", academicYear, academicTerm, userID).
		Find(&courses).Error
	if err != nil {
		return nil, fmt.Errorf("查询课程失败: %w", err)
	}

	var result []model.Course
	for _, course := range courses {
		weekRanges := strings.Split(course.Week, ",")
		for _, weekRange := range weekRanges {
			rangeParts := strings.Split(weekRange, "-")
			if len(rangeParts) != 2 {
				continue
			}

			startWeek, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				continue
			}
			endWeek, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				continue
			}

			if week >= startWeek && week <= endWeek {
				weekdays := strings.Split(course.Weekday, ",")
				sections := strings.Split(course.Section, ",")
				addresses := strings.Split(course.CourseAddr, ",")

				if len(weekdays) != len(sections) || len(sections) != len(addresses) {
					return nil, fmt.Errorf("课程数据不匹配，weekday、section、course_addr 数量不一致")
				}

				for i := 0; i < len(weekdays); i++ {
					newCourse := model.Course{
						UserID:             course.UserID,
						CourseNumber:       course.CourseNumber,
						CourseName:         course.CourseName,
						CourseImg:          course.CourseImg,
						CourseIntroduction: course.CourseIntroduction,
						CourseClass:        course.CourseClass,
						CourseType:         course.CourseType,
						CourseAddr:         addresses[i],
						LecturerProfile:    course.LecturerProfile,
						Credit:             course.Credit,
						AcademicYear:       course.AcademicYear,
						AcademicTerm:       course.AcademicTerm,
						Week:               course.Week,
						Weekday:            weekdays[i],
						Section:            sections[i],
						IsCompleted:        course.IsCompleted,
					}
					result = append(result, newCourse)
				}
			}
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Weekday != result[j].Weekday {
			return result[i].Weekday < result[j].Weekday
		}
		return result[i].Section < result[j].Section
	})

	return result, nil
}

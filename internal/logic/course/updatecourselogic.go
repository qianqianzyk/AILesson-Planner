package course

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"gorm.io/gorm"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourseLogic {
	return &UpdateCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCourseLogic) UpdateCourse(req *types.UpdateCourseReq) (resp *types.UpdateCourseResp, err error) {
	id := req.ID
	courseNumber := req.CourseNumber
	courseName := req.CourseName
	courseImg := req.CourseImg
	courseIntroduction := req.CourseIntroduction
	courseClass := req.CourseClass
	courseType := req.CourseType
	courseAddr := req.CourseAddr
	lecturerProfile := req.LecturerProfile
	credit := req.Credit
	academicYear := req.AcademicYear
	academicTerm := req.AcademicTerm
	week := req.Week
	weekday := req.Weekday
	section := req.Section
	isCompleted := req.IsCompleted

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	course, err := service.GetCourseByID(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrGetCourse, err)
	}
	if course.UserID != int(userID) {
		return nil, utils.AbortWithException(utils.ErrAuthUser, err)
	}

	if courseImg != course.CourseImg {
		service.DeleteObjectByUrlAsync(course.CourseImg)
		course.CourseImg = courseImg
	}
	course.CourseNumber = courseNumber
	course.CourseName = courseName
	course.CourseIntroduction = courseIntroduction
	course.CourseClass = courseClass
	course.CourseType = courseType
	course.CourseAddr = courseAddr
	course.LecturerProfile = lecturerProfile
	course.Credit = credit
	course.AcademicYear = academicYear
	course.AcademicTerm = academicTerm
	course.Week = week
	course.Weekday = weekday
	course.Section = section
	course.IsCompleted = isCompleted
	err = service.UpdateCourse(course)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUpdateCourse, err)
	}

	return &types.UpdateCourseResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}

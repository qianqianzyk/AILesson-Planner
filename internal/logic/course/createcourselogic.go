package course

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCourseLogic {
	return &CreateCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCourseLogic) CreateCourse(req *types.CreateCourseReq) (resp *types.CreateCourseResp, err error) {
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
	classes := req.Classes

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	course := model.Course{
		UserID:             int(userID),
		CourseNumber:       courseNumber,
		CourseName:         courseName,
		CourseImg:          courseImg,
		CourseIntroduction: courseIntroduction,
		CourseClass:        courseClass,
		CourseType:         courseType,
		CourseAddr:         courseAddr,
		LecturerProfile:    lecturerProfile,
		Credit:             credit,
		AcademicYear:       academicYear,
		AcademicTerm:       academicTerm,
		Week:               week,
		Weekday:            weekday,
		Section:            section,
		IsCompleted:        false,
	}
	err = service.CreateCourse(&course)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrCreateCourse, err)
	}

	if len(classes) != 0 {
		err = service.CreateScoresBatch(int(course.ID), classes)
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrServer, err)
		}
	}

	return &types.CreateCourseResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}

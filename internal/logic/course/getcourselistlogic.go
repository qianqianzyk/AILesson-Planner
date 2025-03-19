package course

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCourseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseListLogic {
	return &GetCourseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseListLogic) GetCourseList(req *types.Empty) (resp *types.GetCourseListResp, err error) {
	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	courses, err := service.GetCourseList(int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	var courseList []types.Course
	for _, c := range courses {
		classList, err := service.GetClassListByCourseID(int(c.ID))
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrGetClassList, err)
		}

		courseList = append(courseList, types.Course{
			ID:                 int(c.ID),
			CourseNumber:       c.CourseNumber,
			CourseName:         c.CourseName,
			CourseImg:          c.CourseImg,
			CourseIntroduction: c.CourseIntroduction,
			CourseClass:        c.CourseClass,
			CourseType:         c.CourseType,
			CourseAddr:         c.CourseAddr,
			LecturerProfile:    c.LecturerProfile,
			Credit:             c.Credit,
			AcademicYear:       c.AcademicYear,
			AcademicTerm:       c.AcademicTerm,
			Week:               c.Week,
			Weekday:            c.Weekday,
			Section:            c.Section,
			IsCompleted:        c.IsCompleted,
			Classes:            classList,
		})
	}

	return &types.GetCourseListResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: courseList,
	}, nil
}

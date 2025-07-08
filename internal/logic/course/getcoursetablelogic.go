package course

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseTableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCourseTableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseTableLogic {
	return &GetCourseTableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseTableLogic) GetCourseTable(req *types.GetCourseTableReq) (resp *types.GetCourseTableResp, err error) {
	academicYear := req.AcademicYear
	academicTerm := req.AcademicTerm
	week := req.Week

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	courses, err := service.GetCoursesByWeek(academicYear, academicTerm, int(userID), week)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	var coursesTable []types.CourseTable
	for _, c := range courses {
		coursesTable = append(coursesTable, types.CourseTable{
			CourseName:  c.CourseName,
			CourseClass: c.CourseClass,
			CourseType:  c.CourseType,
			CourseAddr:  c.CourseAddr,
			Credit:      c.Credit,
			Week:        c.Week,
			Weekday:     c.Weekday,
			Section:     c.Section,
		})
	}

	return &types.GetCourseTableResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: coursesTable,
	}, nil
}

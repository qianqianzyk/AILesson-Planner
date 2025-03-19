package score

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

type ExportScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportScoresLogic {
	return &ExportScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportScoresLogic) ExportScores(req *types.ExportScoresReq) (resp *types.ExportScoresResp, err error) {
	courseID := req.CourseID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	course, err := service.GetCourseByID(courseID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrGetCourse, err)
	}
	if course.UserID != int(userID) {
		return nil, utils.AbortWithException(utils.ErrAuthUser, err)
	}

	courses, err := service.GetStudentScore(courseID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetStudentScore, err)
	}

	url, err := service.ExportStudentScores(course.CourseClass, courses)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrExportScores, err)
	}

	return &types.ExportScoresResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.FileUrl{
			Url: url,
		},
	}, nil
}

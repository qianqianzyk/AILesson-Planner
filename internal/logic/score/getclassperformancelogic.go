package score

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassPerformanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassPerformanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassPerformanceLogic {
	return &GetClassPerformanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassPerformanceLogic) GetClassPerformance(req *types.GetClassPerformanceReq) (resp *types.GetClassPerformanceResp, err error) {
	class := req.Class
	year := req.Year
	term := req.Term

	courseAvgScores, classAvgGradePoint, err := service.GetClassPerformance(year, term, class)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	var classAvgScores []types.CourseAvgScore
	for _, c := range courseAvgScores {
		classAvgScores = append(classAvgScores, types.CourseAvgScore{
			CourseID:   c.CourseID,
			CourseName: c.CourseName,
			AvgScore:   c.AvgScore,
		})
	}

	return &types.GetClassPerformanceResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ClassPerformance{
			ClassAvgGradePoint: classAvgGradePoint.AvgGradePoint,
			CourseAvgScores:    classAvgScores,
		},
	}, nil
}

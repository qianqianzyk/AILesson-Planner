package score

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCountNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCountNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCountNumberLogic {
	return &GetCountNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCountNumberLogic) GetCountNumber(req *types.GetCountPerformanceNumberReq) (resp *types.GetCountPerformanceNumberResp, err error) {
	class := req.Class
	year := req.Year
	term := req.Term
	courseName := req.CourseName

	overView, err := service.GetScoreOverview(year, term, class, courseName)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	return &types.GetCountPerformanceNumberResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.CountPerformanceNumber{
			ExcellentCount: overView.ExcellentCount,
			GoodCount:      overView.GoodCount,
			PassCount:      overView.PassCount,
			FailCount:      overView.FailCount,
		},
	}, nil
}

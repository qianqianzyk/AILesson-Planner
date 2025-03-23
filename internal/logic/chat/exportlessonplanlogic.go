package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportLessonPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportLessonPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportLessonPlanLogic {
	return &ExportLessonPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportLessonPlanLogic) ExportLessonPlan(req *types.ExportLessonPlanReq) (resp *types.ExportLessonPlanResp, err error) {
	id := req.ID

	tPlan, err := service.GetTPlanByID(id)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	tPlanUrl, err := service.GenerateWordDoc(tPlan)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrExportTPlan, err)
	}

	return &types.ExportLessonPlanResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.FileUrl{
			Url: tPlanUrl,
		},
	}, nil
}

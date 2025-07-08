package share

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseNameListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCourseNameListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseNameListLogic {
	return &GetCourseNameListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseNameListLogic) GetCourseNameList(req *types.Empty) (resp *types.GetCourseNameListResp, err error) {
	return &types.GetCourseNameListResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: service.GetAllSubjects(),
	}, nil
}
